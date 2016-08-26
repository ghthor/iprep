package main

import (
	"fmt"
	"log"
	"strings"
)

const (
	BracesOpening = "{[("
	BracesClosing = "}])"
)

type BraceStack []rune

func (s BraceStack) Push(r rune) BraceStack {
	return append(s, r)
}

func (s BraceStack) Top() rune {
	if len(s) == 0 {
		return 0
	}

	return s[len(s)-1]
}

func ValidateBraces(src string) bool {
	braces := make(BraceStack, 0, len(src))

	for _, r := range src {
		// TODO: Fix Dangerous handling of UTF8 and string byte indexs
		if idx := strings.IndexRune(BracesOpening, r); idx >= 0 {
			braces = braces.Push(r)
			continue
		}

		if idx := strings.IndexRune(BracesClosing, r); idx >= 0 {
			opening, _, err := strings.NewReader(BracesOpening[idx:]).ReadRune()
			if err != nil {
				log.Fatal(err)
			}

			if opening == braces.Top() {
				braces = braces[:len(braces)-1]
				continue
			}

			return false
		}
	}

	return true
}

func main() {
	fmt.Println(ValidateBraces("()"))
	fmt.Println(ValidateBraces("[()]"))
	fmt.Println(ValidateBraces("{(){}]"))
}
