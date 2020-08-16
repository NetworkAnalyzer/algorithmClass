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

// count of brackets Brackets
type Brackets struct {
	state struct {
		round  int
		curly  int
		square int
	}
	banned []string
}

func isOpen(bracket string) bool {
	return funk.Contains([]string{RoundOpen, CurlyOpen, SquareOpen}, bracket)
}

func isClose(bracket string) bool {
	return funk.Contains([]string{RoundClose, CurlyClose, SquareClose}, bracket)
}

func (b *Brackets) canClose(bracket string) bool {
	return !funk.Contains(b.banned, bracket)
}

func (b *Brackets) isAllClosed() bool {
	return b.state.round == 0 && b.state.curly == 0 && b.state.square == 0
}

func (b *Brackets) open(bracket string) {
	switch bracket {
	case RoundOpen:
		b.state.round++
		b.banned = []string{CurlyClose, SquareClose}
	case CurlyOpen:
		b.state.curly++
		b.banned = []string{RoundClose, SquareClose}
	case SquareOpen:
		b.state.square++
		b.banned = []string{RoundClose, CurlyClose}
	}
}

func (b *Brackets) close(bracket string) bool {
	if !b.canClose(bracket) {
		return false
	}

	switch bracket {
	case RoundClose:
		b.state.round--
	case CurlyClose:
		b.state.curly--
	case SquareClose:
		b.state.square--
	}

	b.resetBanned()
	return true
}

func (b *Brackets) resetBanned() {
	b.banned = []string{}
}

func Run(str string) (bool, error) {
	r := regexp.MustCompile(`^[\(\)\[\]\{\}]+$`)
	if !r.MatchString(str) {
		return false, errors.New("allowed only brackets")
	}

	chars := strings.Split(str, "")
	brackets := Brackets{}

	for _, char := range chars {
		if isOpen(char) {
			brackets.open(char)
		} else if isClose(char) {
			if ok := brackets.close(char); !ok {
				return false, nil
			}
		}
	}

	return brackets.isAllClosed(), nil
}
