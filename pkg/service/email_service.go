package service

import (
	"errors"
	"strconv"
	"strings"
)

type service struct {
}

func NewEmailService() *service {
	return &service{}
}

func (s *service) EmailCheck(m string) (string, error) {
	res := strings.TrimSpace(m)
	res1 := strings.TrimPrefix(res, "\n")
	if !CheckEmail(res1) {
		return "", errors.New("bad req")
	}
	return res1, nil
}
func CheckEmail(s string) bool {
	for _, i := range s {
		if i == '@' {
			for _, j := range s {
				if j == '.' && j != rune(s[len(s)-1]) {

					return true
				}
			}

		}
	}

	return false
}

func (s *service) IinCheck(m string) (string, error) {
	res := strings.TrimSpace(string(m))
	res1 := strings.TrimPrefix(res, "\n")
	if !CheckIin(res1) || 11 != len(m)-1 {
		return "", errors.New("bad req")
	}
	return res1, nil
}
func CheckIin(s string) bool {

	first := s[:2]
	if first[0] == '0' && first[1] >= '6' {

		return false
	}
	second := s[2:4]
	secondY, _ := strconv.Atoi(second)
	if secondY > 12 {
		return false
	}
	days := s[4:6]
	daysY, _ := strconv.Atoi(days)
	if daysY > 31 {
		return false
	}

	return true
}
