package main

import (
	"errors"
	"strings"
)

type StringService interface {
	UpperCase(string) (string, error)
	Count(string) int
}

type stringService struct{}

var errEmpty = errors.New("emtpy string")

func (stringService) UpperCase(str string) (string, error) {
	if str == "" {
		return "", errEmpty
	}
	return strings.ToUpper(str), nil
}

func (stringService) Count(str string) int {
	return len(str)
}
