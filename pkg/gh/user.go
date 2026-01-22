package gh

import (
	"context"

	"github.com/google/go-github/v60/github"
)

// GetUser fetches a GitHub user by their username.
func (c *Client) GetUser(username string) (*github.User, error) {
	ctx := context.Background()
	user, _, err := c.Users.Get(ctx, username)
	return user, err
}
