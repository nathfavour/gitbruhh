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
	return nil, fmt.Errorf("scraping search not implemented yet")
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
