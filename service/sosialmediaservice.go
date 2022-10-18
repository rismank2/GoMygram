package service

import "errors"

type SocialMediaInterf interface {
	CekPostSocialMedia(name, social_media_url string) error
}
type SocialmediaSer struct{}

// CekInputanSocialMedia implements SocialmediaIfac
func (sc *SocialmediaSer) CekPostSocialMedia(name, social_media_url string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	if social_media_url == "" {
		return errors.New("social_media_url cannot be empty")
	}
	return nil
}

func SocialMediaSerN() SocialMediaInterf {
	return &SocialmediaSer{}
}
