package gh

import (
	"context"
	"os"

	"github.com/google/go-github/v60/github"
	"golang.org/x/oauth2"
)

// Client wraps the GitHub client and provides utility methods.
type Client struct {
	*github.Client
	scraper      *Scraper
	providerType ProviderType
}

// NewClient initializes a new GitHub client with a specific provider type.
func NewClient(pType ProviderType) *Client {
	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")

	var ghClient *github.Client
	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		ghClient = github.NewClient(tc)
	} else {
		// Unauthenticated client
		ghClient = github.NewClient(nil)
	}

	return &Client{
		Client:       ghClient,
		scraper:      NewScraper(),
		providerType: pType,
	}
}
