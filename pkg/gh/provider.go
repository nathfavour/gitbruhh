package gh

import (
	"github.com/google/go-github/v60/github"
)

// ProviderType defines the type of provider to use.
type ProviderType string

const (
	ProviderAuto    ProviderType = "auto"
	ProviderAPI     ProviderType = "api"
	ProviderScraper ProviderType = "scraper"
)

// Provider defines the interface for fetching GitHub data.
type Provider interface {
	GetRepo(fullName string) (*github.Repository, error)
	GetIssues(fullName string) ([]*github.Issue, error)
	GetUser(username string) (*github.User, error)
	SearchRepos(query string) ([]*github.Repository, error)
}
