package utils

import (
	"errors"
	"regexp"
)

func ArePasswordsSimilar(password, confirmPassword string) error {
	if password != confirmPassword {
		return errors.New("passwords don't match")
	}
	return nil
}

func RegexCheck(input, pattern string) error {
	match, _ := regexp.MatchString(pattern, input)
	if !match {
		return errors.New("password must contain at least 1 lowercase letter, 1 uppercase letter, 1 digit, and 1 special character. It must be at least 8 characters long")
	}
	return nil
}
