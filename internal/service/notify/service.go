package notify

import (
	"fmt"

	"go-atch-template/internal/model"
)

type sender interface {
	Send(login, message string) error
}

type service struct {
	s sender
}

func New(s sender) service {
	return service{s: s}
}

func (s service) NotifyUser(user model.User) error {
	if !user.EnableNotify {
		return nil
	}

	err := s.s.Send(user.Name, "register detected")
	if err != nil {
		return fmt.Errorf("faile to send: %w", err)
	}
	return nil
}
