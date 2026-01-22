package gh

import (
	"fmt"
	"net/http"
	"io"
	"strings"
	"regexp"
	"strconv"

	"github.com/google/go-github/v60/github"
)

// Scraper implements the Provider interface by scraping GitHub web pages.
type Scraper struct{}

func NewScraper() *Scraper {
	return &Scraper{}
}

func (s *Scraper) GetRepo(fullName string) (*github.Repository, error) {
	url := fmt.Sprintf("https://github.com/%s", fullName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch page: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	html := string(body)
	repo := &github.Repository{
		FullName: github.String(fullName),
	}

	// Very basic scraping examples
	// Extract description
	reDesc := regexp.MustCompile(`<p class="f4 my-3">\s*(.*?)\s*</p>`)
	matchDesc := reDesc.FindStringSubmatch(html)
	if len(matchDesc) > 1 {
		repo.Description = github.String(strings.TrimSpace(matchDesc[1]))
	}

	// Extract stars
	reStars := regexp.MustCompile(`id="repo-stars-counter-star" .*?>(.*?)</span>`)
	matchStars := reStars.FindStringSubmatch(html)
	if len(matchStars) > 1 {
		starsStr := strings.TrimSpace(matchStars[1])
		// Handle 1.2k etc
		stars := parseGitHubCount(starsStr)
		repo.StargazersCount = github.Int(stars)
	}

	// Extract forks
	reForks := regexp.MustCompile(`id="repo-network-counter" .*?>(.*?)</span>`)
	matchForks := reForks.FindStringSubmatch(html)
	if len(matchForks) > 1 {
		repo.ForksCount = github.Int(parseGitHubCount(matchForks[1]))
	}

	// Extract issues count
	reIssues := regexp.MustCompile(`id="issues-repo-tab" .*?><span .*?>Issues</span><span .*?>(.*?)</span>`)
	matchIssues := reIssues.FindStringSubmatch(html)
	if len(matchIssues) > 1 {
		repo.OpenIssuesCount = github.Int(parseGitHubCount(matchIssues[1]))
	}

	// Extract language
	reLang := regexp.MustCompile(`itemprop="programmingLanguage">(.*?)</span>`)
	matchLang := reLang.FindStringSubmatch(html)
	if len(matchLang) > 1 {
		repo.Language = github.String(strings.TrimSpace(matchLang[1]))
	}

	// Extract license
	reLicense := regexp.MustCompile(`itemprop="license">(.*?)</a>`)
	if matchLicense := reLicense.FindStringSubmatch(html); len(matchLicense) > 1 {
		repo.License = &github.License{Name: github.String(strings.TrimSpace(matchLicense[1]))}
	} else {
		// Try alternative license pattern
		reLicenseAlt := regexp.MustCompile(`<svg .*? octicon-law .*?>.*?</svg>\s*(.*?)\s*</a>`)
		if matchLicenseAlt := reLicenseAlt.FindStringSubmatch(html); len(matchLicenseAlt) > 1 {
			repo.License = &github.License{Name: github.String(strings.TrimSpace(matchLicenseAlt[1]))}
		}
	}

	repo.HTMLURL = github.String(url)

	return repo, nil
}

func (s *Scraper) GetIssues(fullName string) ([]*github.Issue, error) {
	// Basic implementation for now
	return nil, fmt.Errorf("scraping issues not implemented yet")
}

func (s *Scraper) GetUser(username string) (*github.User, error) {
	url := fmt.Sprintf("https://github.com/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch page: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	html := string(body)
	user := &github.User{
		Login: github.String(username),
	}

	// Extract Name
	reName := regexp.MustCompile(`<span class="p-name vcard-fullname d-block overflow-hidden" itemprop="name">\s*(.*?)\s*</span>`)
	matchName := reName.FindStringSubmatch(html)
	if len(matchName) > 1 {
		user.Name = github.String(strings.TrimSpace(matchName[1]))
	}

	user.HTMLURL = github.String(url)

	return user, nil
}

func (s *Scraper) SearchRepos(query string) ([]*github.Repository, error) {
	url := fmt.Sprintf("https://github.com/search?q=%s&type=repositories", strings.ReplaceAll(query, " ", "+"))
	req, _ := http.NewRequest("GET", url, nil)
	// Add user agent to look like a browser
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch search results: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	html := string(body)
	var repos []*github.Repository

	// Search results on GitHub web are a bit complex to parse with regex
	// But we can try to find repo names which are usually in <a> tags with data-hydro-click
	reRepo := regexp.MustCompile(`href="/([a-zA-Z0-9-._]+/[a-zA-Z0-9-._]+)" data-hydro-click`)
	matches := reRepo.FindAllStringSubmatch(html, 10)

	for _, match := range matches {
		fullName := match[1]
		repos = append(repos, &github.Repository{
			FullName: github.String(fullName),
			HTMLURL:  github.String("https://github.com/" + fullName),
		})
	}

	if len(repos) == 0 {
		// Fallback regex for search results
		reRepoAlt := regexp.MustCompile(`"repository":\{"id":\d+,"name":"(.*?)","owner":\{"login":"(.*?)"\}`)
		matchesAlt := reRepoAlt.FindAllStringSubmatch(html, 10)
		for _, match := range matchesAlt {
			fullName := match[2] + "/" + match[1]
			repos = append(repos, &github.Repository{
				FullName: github.String(fullName),
				HTMLURL:  github.String("https://github.com/" + fullName),
			})
		}
	}

	return repos, nil
}

func parseGitHubCount(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	multiplier := 1
	if strings.HasSuffix(s, "k") {
		multiplier = 1000
		s = strings.TrimSuffix(s, "k")
	} else if strings.HasSuffix(s, "m") {
		multiplier = 1000000
		s = strings.TrimSuffix(s, "m")
	}
	
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return int(val * float64(multiplier))
}
