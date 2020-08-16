package problem020

import (
	"errors"
	"github.com/thoas/go-funk"
	"regexp"
	"strings"
)

// This problem was asked by Facebook.
//
// Given a string of round, curly, and square open and closing Brackets, return whether the Brackets are balanced (well-formed).
// For example, given the string "([])[]({})", you should return true.
// Given the string "([)]" or "((()", you should return false.

const (
	RoundOpen   = "("
	RoundClose  = ")"
	CurlyOpen   = "{"
	CurlyClose  = "}"
	SquareOpen  = "["
	SquareClose = "]"
)

var bracketPair = map[string]string{
	RoundClose:  RoundOpen,
	CurlyClose:  CurlyOpen,
	SquareClose: SquareOpen,
}

func Run(str string) (bool, error) {
	r := regexp.MustCompile(`^[\(\)\[\]\{\}]+$`)
	if !r.MatchString(str) {
		return false, errors.New("allowed only brackets")
	}

	stack := make([]string, 0)

	for _, char := range strings.Split(str, "") {
		if isOpen(char) {
			stack = append(stack, char)
		} else {
			if last(stack) == bracketPair[char] {
				stack = removeLast(stack)
			}
		}
	}

	return len(stack) == 0, nil
}

// utils
func isOpen(bracket string) bool {
	return funk.Contains([]string{RoundOpen, CurlyOpen, SquareOpen}, bracket)
}

func last(slice []string) string {
	return slice[len(slice) - 1]
}

func removeLast(slice []string) []string {
	return slice[:len(slice) - 1]
}
