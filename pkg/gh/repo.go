package gh

import (
	"context"
	"errors"
	"strings"

	"github.com/google/go-github/v60/github"
)

// GetRepo fetches a repository by its owner and name (e.g., "owner/repo").
func (c *Client) GetRepo(fullName string) (*github.Repository, error) {
	ctx := context.Background()
	parts := strings.Split(fullName, "/")
	if len(parts) != 2 {
		return nil, errors.New("invalid repo format; use 'owner/repo'")
	}
	repo, _, err := c.Repositories.Get(ctx, parts[0], parts[1])
	return repo, err
}

// GetIssues fetches the most recent open issues for a repository.
func (c *Client) GetIssues(fullName string) ([]*github.Issue, error) {
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
	return issues, err
}
