package service

import "errors"

type Photointerf interface {
	CekPostPhoto(title, url string) error
}

type PhotoService struct {
}

// CekPostPhoto implements Photointerf
func (ps *PhotoService) CekPostPhoto(title string, url string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}
	if url == "" {
		return errors.New("url cannot be empty")
	}
	return nil
}

func NewPhotoService() Photointerf {
	return &PhotoService{}
}
