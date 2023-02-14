package models

import (
	// We use the "testify" library for mocking our store
	"context"
	"testing"

	"github.com/google/go-github/v48/github"
	"github.com/stretchr/testify/assert"
)

func TestGitHubSuccess(t *testing.T) {

	context := context.TODO()

	loginName := "justAUser"
	user := github.User{

		Login: &loginName,
	}

	m := new(MockSocialMedia)
	m.On("RetrieveUserInfoGitHub", context).Return(&user, nil)

	s := SocialMediaSignup{m}

	result, err := s.Signup("GitHub", context)

	m.AssertExpectations(t)
	// Finally, we assert that we should'nt get any error
	if err != nil {
		t.Errorf("error should be nil, got: %v", err)
	}

	assert.Equal(t, *(result.(*github.User)).Login, loginName, "The two words are the same")

}
