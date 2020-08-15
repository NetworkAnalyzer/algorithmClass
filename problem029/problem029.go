package problem029

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// This problem was asked by Amazon.
//
// Run-length encoding is a fast and simple method of encoding strings.
// The basic idea is to represent repeated successive characters as a single count and character.
//
// For example, the string "AAAABBBCCDAA" would be encoded as "4A3B2C1D2A".
//
// Implement run-length encoding and decoding.
// You can assume the string to be encoded have no digits and consists solely of alphabetic characters.
// You can assume the string to be decoded is valid.

type previousCharInfo struct {
	char  string
	count int
}

func (p *previousCharInfo) encode() string {
	return fmt.Sprintf("%d%s", p.count, p.char)
}

func (p *previousCharInfo) set(char string) {
	p.char = char
}

func (p *previousCharInfo) countUp() {
	p.count++
}

func (p *previousCharInfo) resetCount() {
	p.count = 1
}

func encode(str string) (string, error) {
	r := regexp.MustCompile(`^[A-Z]+$`)
	if !r.MatchString(str) {
		return "", errors.New("format error")
	}

	var result string
	var previous previousCharInfo
	for i, v := range str {
		char := string(v)

		if i == 0 || char == previous.char {
			previous.set(char)
			previous.countUp()
			continue
		}

		result += previous.encode()
		previous.set(char)
		previous.resetCount()
	}

	result += previous.encode()

	return result, nil
}

func decode(str string) (string, error) {
	r := regexp.MustCompile(`^(\d[A-Z])+$`)
	if !r.MatchString(str) {
		return "", errors.New("format error")
	}

	var result string
	chars := strings.Split(str, "")
	for i := 0; i < len(chars); i += 2 {
		if repeat, err := strconv.Atoi(chars[i]); err != nil {
			return "", err
		} else {
			result += strings.Repeat(chars[i + 1], repeat)
		}
	}

	return result, nil
}
