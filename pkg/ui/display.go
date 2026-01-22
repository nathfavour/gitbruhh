package ui

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/google/go-github/v60/github"
)

// PrintUser displays user information in a clean table format.
func PrintUser(user *github.User) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "FIELD\tVALUE")
	fmt.Fprintf(w, "Login\t%s\n", user.GetLogin())
	fmt.Fprintf(w, "Name\t%s\n", user.GetName())
	fmt.Fprintf(w, "Bio\t%s\n", user.GetBio())
	fmt.Fprintf(w, "Location\t%s\n", user.GetLocation())
	fmt.Fprintf(w, "Followers\t%d\n", user.GetFollowers())
	fmt.Fprintf(w, "Following\t%d\n", user.GetFollowing())
	fmt.Fprintf(w, "Public Repos\t%d\n", user.GetPublicRepos())
	fmt.Fprintf(w, "URL\t%s\n", user.GetHTMLURL())
	w.Flush()
}

// PrintRepo displays repository information in a clean table format.
func PrintRepo(repo *github.Repository) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "FIELD\tVALUE")
	fmt.Fprintf(w, "Full Name\t%s\n", repo.GetFullName())
	fmt.Fprintf(w, "Description\t%s\n", repo.GetDescription())
	fmt.Fprintf(w, "Language\t%s\n", repo.GetLanguage())
	fmt.Fprintf(w, "Stars\t%d\n", repo.GetStargazersCount())
	fmt.Fprintf(w, "Forks\t%d\n", repo.GetForksCount())
	fmt.Fprintf(w, "Issues\t%d\n", repo.GetOpenIssuesCount())
	fmt.Fprintf(w, "License\t%s\n", repo.GetLicense().GetName())
	fmt.Fprintf(w, "URL\t%s\n", repo.GetHTMLURL())
	w.Flush()
}

// PrintRepos displays a list of repositories in a clean table format.
func PrintRepos(repos []*github.Repository) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "NAME\tSTARS\tLANGUAGE\tDESCRIPTION")
	for _, repo := range repos {
		desc := repo.GetDescription()
		if len(desc) > 50 {
			desc = desc[:47] + "..."
		}
		fmt.Fprintf(w, "%s\t%d\t%s\t%s\n",
			repo.GetFullName(),
			repo.GetStargazersCount(),
			repo.GetLanguage(),
			desc,
		)
	}
	w.Flush()
}

// PrintIssues displays a list of issues in a clean table format.
func PrintIssues(issues []*github.Issue) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "NUMBER\tTITLE\tUSER\tURL")
	for _, issue := range issues {
		title := issue.GetTitle()
		if len(title) > 50 {
			title = title[:47] + "..."
		}
		fmt.Fprintf(w, "#%d\t%s\t%s\t%s\n",
			issue.GetNumber(),
			title,
			issue.GetUser().GetLogin(),
			issue.GetHTMLURL(),
		)
	}
	w.Flush()
}

