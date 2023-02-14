package models

import (
	"context"
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/google/go-github/v48/github"
	google "google.golang.org/api/oauth2/v2"
)

// unexported key type prevents collisions
type key int

const (
	userKey key = iota
)

type SocialMedia interface {
	RetrieveUserInfoGitHub(ctx context.Context) (*github.User, error)
	RetrieveUserInfoGoogle(ctx context.Context) (*google.Userinfo, error)
	RetrieveUserInfoTwitter(ctx context.Context) (*twitter.User, error)
}

// WithUser returns a copy of ctx that stores the Github User.
func WithUserGitHub(ctx context.Context, user *github.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func RetrieveUserInfoGitHub(ctx context.Context) (*github.User, error) {
	user, ok := ctx.Value(userKey).(*github.User)
	if !ok {
		return nil, fmt.Errorf("github: Context missing Github User")
	}
	return user, nil
}

func RetrieveUserInfoGoogle(ctx context.Context) (*google.Userinfo, error) {
	user, ok := ctx.Value(userKey).(*google.Userinfo)
	if !ok {
		return nil, fmt.Errorf("github: Context missing Github User")
	}
	return user, nil
}

func RetrieveUserInfoTwitter(ctx context.Context) (*twitter.User, error) {
	user, ok := ctx.Value(userKey).(*twitter.User)
	if !ok {
		return nil, fmt.Errorf("github: Context missing Github User")
	}
	return user, nil
}
