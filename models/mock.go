package models

import (
	"context"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/google/go-github/v48/github"
	"github.com/stretchr/testify/mock"
	google "google.golang.org/api/oauth2/v2"
)

// Create a MockStore struct with an embedded mock instance
type MockSocialMedia struct {
	mock.Mock
}

func (m *MockSocialMedia) RetrieveUserInfoGitHub(ctx context.Context) (*github.User, error) {
	// This allows us to pass in mocked results, so that the mock store will return whatever we define
	returnVals := m.Called(ctx)

	// return the values which we define
	return returnVals.Get(0).(*github.User), returnVals.Error(1)
}

func (m *MockSocialMedia) RetrieveUserInfoGoogle(ctx context.Context) (*google.Userinfo, error) {
	// This allows us to pass in mocked results, so that the mock store will return whatever we define
	returnVals := m.Called(ctx)

	// return the values which we define
	return returnVals.Get(0).(*google.Userinfo), returnVals.Error(1)
}

func (m *MockSocialMedia) RetrieveUserInfoTwitter(ctx context.Context) (*twitter.User, error) {
	// This allows us to pass in mocked results, so that the mock store will return whatever we define
	returnVals := m.Called(ctx)

	// return the values which we define
	return returnVals.Get(0).(*twitter.User), returnVals.Error(1)
}
