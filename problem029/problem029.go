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

func encode(str string) (string, error) {
	r := regexp.MustCompile(`^[A-Z]+$`)
	if !r.MatchString(str) {
		return "", errors.New("format error")
	}

	var result string
	var previous struct{
		rune  string
		count int
	}
	for i, rune := range str {
		if i == 0 || string(rune) == previous.rune {
			previous.rune = string(rune)
			previous.count++
			continue
		}

		result += fmt.Sprintf("%d%s", previous.count, previous.rune)
		previous.rune = string(rune)
		previous.count = 1
	}

	result += fmt.Sprintf("%d%s", previous.count, previous.rune)

	return result, nil
}

func decode(str string) (string, error) {
	r := regexp.MustCompile(`^(\d[A-Z])+$`)
	if !r.MatchString(str) {
		return "", errors.New("format error")
	}

	var result string
	runes := strings.Split(str, "")
	for i := 0; i < len(runes); i+=2 {
		repeat, err := strconv.Atoi(runes[i])
		if err != nil {
			return "", err
		}
		result += strings.Repeat(runes[i + 1], repeat)
	}

	return result, nil
}
