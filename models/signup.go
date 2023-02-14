package models

import "context"

type SocialMediaSignup struct {
	SocialMedia SocialMedia
}

func (s *SocialMediaSignup) Signup(signUpWay string, ctx context.Context) (interface{}, error) {

	if signUpWay == "GitHub" {

		result, err := s.SocialMedia.RetrieveUserInfoGitHub(ctx)
		if err != nil {

			return nil, err
		} else {
			return result, nil
		}

	} else if signUpWay == "Google" {
		result, err := s.SocialMedia.RetrieveUserInfoGoogle(ctx)
		if err != nil {

			return nil, err
		} else {
			return result, nil
		}

	} else if signUpWay == "Twitter" {

		result, err := s.SocialMedia.RetrieveUserInfoTwitter(ctx)
		if err != nil {

			return nil, err
		} else {
			return result, nil
		}

	}

	return nil, nil

}
