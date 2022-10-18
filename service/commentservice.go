package service

import "errors"

type Commentinterf interface {
	CekPostComment(message string) error
}

type CommentService struct{}

func CommentServic() Commentinterf {
	return &CommentService{}
}

func (cs *CommentService) CekPostComment(message string) error {
	if message == "" {
		return errors.New("message cannot be empty")
	}
	return nil
}
