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
}

// NewClient initializes a new GitHub client.
// It looks for GITHUB_TOKEN in the environment for authentication.
func NewClient() *Client {
	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")

	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		return &Client{github.NewClient(tc)}
	}

	// Unauthenticated client
	return &Client{github.NewClient(nil)}
}
