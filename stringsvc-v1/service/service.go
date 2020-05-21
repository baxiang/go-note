package service

import (
	"errors"
	"strings"
)


type ServiceMiddleware func(StringService)StringService

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}
type StringServiceImpl struct{}

func (StringServiceImpl) Uppercase(s string) (string, error) {
	if s == "" {
		return "", errors.New("empty string")
	}
	return strings.ToUpper(s), nil
}

func (StringServiceImpl) Count(s string) int {
	return len(s)
}
