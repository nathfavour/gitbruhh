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
