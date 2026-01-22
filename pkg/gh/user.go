package gh

import (
	"context"

	"github.com/google/go-github/v60/github"
)

// GetUser fetches a GitHub user by their username.
func (c *Client) GetUser(username string) (*github.User, error) {
	if c.providerType == ProviderScraper {
		return c.scraper.GetUser(username)
	}

	ctx := context.Background()
	user, _, err := c.Users.Get(ctx, username)
	if err != nil && (c.providerType == ProviderAuto || c.providerType == "") {
		// Fallback to scraper
		if scrapedUser, sErr := c.scraper.GetUser(username); sErr == nil {
			return scrapedUser, nil
		}
	}
	return user, err
}
