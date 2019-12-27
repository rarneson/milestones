package github

import (
	"context"

	gitHubClient "github.com/google/go-github/v27/github"
	"golang.org/x/oauth2"
)

// GitHub stores our context and client
type GitHub struct {
	ctx    context.Context
	client *gitHubClient.Client
	Org    string
}

// Options stores values needed to interact with the GitHub API
type Options struct {
	Token  string
	Server string
}

// New creates new GitHub instance
func New(options Options) (*GitHub, error) {
	ctx := context.Background()
	tokenservice := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: options.Token})
	tokenclient := oauth2.NewClient(ctx, tokenservice)
	api := options.Server + "api/v3"

	client, err := gitHubClient.NewEnterpriseClient(api, api, tokenclient)
	if err != nil {
		return nil, err
	}

	return &GitHub{
		ctx:    ctx,
		client: client,
	}, nil
}
