package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func ArePasswordsSimilar(password, confirmPassword string) error {
	if password != confirmPassword {
		return errors.New("passwords don't match")
	}
	return nil
}

func RegexCheck(input, pattern string) error {

	// Compile the regex pattern
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return err
	}

	// Match the string
	match := re.MatchString(input)

	fmt.Println(input, pattern)
	fmt.Println(match)
	if !match {
		return errors.New("password must contain at least 1 lowercase letter, 1 uppercase letter, 1 digit, and 1 special character. It must be at least 8 characters long")
	}
	var hasLower, hasUpper, hasDigit, hasSpecial bool
	for _, char := range input {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasDigit = true
		case strings.ContainsRune("@$!%*?&", char):
			hasSpecial = true
		}
	}
	if !(hasLower && hasUpper && hasDigit && hasSpecial) {
		return errors.New("password must contain at least 1 lowercase letter, 1 uppercase letter, 1 digit, and 1 special character. It must be at least 8 characters long")

	}

	return nil
}
