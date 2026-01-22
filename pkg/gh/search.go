package gh

import (
	"context"

	"github.com/google/go-github/v60/github"
)

// SearchRepos searches for repositories matching the query.
func (c *Client) SearchRepos(query string) ([]*github.Repository, error) {
	if c.providerType == ProviderScraper {
		return c.scraper.SearchRepos(query)
	}

	ctx := context.Background()
	opts := &github.SearchOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}
	result, _, err := c.Search.Repositories(ctx, query, opts)
	if err != nil && (c.providerType == ProviderAuto || c.providerType == "") {
		// Fallback to scraper
		if scrapedRepos, sErr := c.scraper.SearchRepos(query); sErr == nil {
			return scrapedRepos, nil
		}
	}
	return result.Repositories, nil
}
