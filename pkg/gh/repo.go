package gh

import (
	"context"
	"errors"
	"strings"

	"github.com/google/go-github/v60/github"
)

// GetRepo fetches a repository by its owner and name (e.g., "owner/repo").
func (c *Client) GetRepo(fullName string) (*github.Repository, error) {
	if c.providerType == ProviderScraper {
		return c.scraper.GetRepo(fullName)
	}

	ctx := context.Background()
	parts := strings.Split(fullName, "/")
	if len(parts) != 2 {
		return nil, errors.New("invalid repo format; use 'owner/repo'")
	}
	repo, _, err := c.Repositories.Get(ctx, parts[0], parts[1])
	if err != nil && (c.providerType == ProviderAuto || c.providerType == "") {
		// Fallback to scraper if API fails
		if scrapedRepo, sErr := c.scraper.GetRepo(fullName); sErr == nil {
			return scrapedRepo, nil
		}
	}
	return repo, err
}

// GetIssues fetches the most recent open issues for a repository.
func (c *Client) GetIssues(fullName string) ([]*github.Issue, error) {
	if c.providerType == ProviderScraper {
		return c.scraper.GetIssues(fullName)
	}

	ctx := context.Background()
	parts := strings.Split(fullName, "/")
	if len(parts) != 2 {
		return nil, errors.New("invalid repo format; use 'owner/repo'")
	}
	opts := &github.IssueListByRepoOptions{
		State: "open",
		ListOptions: github.ListOptions{PerPage: 10},
	}
	issues, _, err := c.Issues.ListByRepo(ctx, parts[0], parts[1], opts)
	if err != nil && (c.providerType == ProviderAuto || c.providerType == "") {
		// Fallback to scraper if API fails
		if scrapedIssues, sErr := c.scraper.GetIssues(fullName); sErr == nil {
			return scrapedIssues, nil
		}
	}
	return issues, err
}
