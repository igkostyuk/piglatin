package main

import (
	"strings"
	"unicode"
)

const (
	vowelSuffix     = "yay"
	consonantSuffix = "ay"
)

func lower(r rune) rune {
	if 'A' <= r && r <= 'Z' {
		r += 'a' - 'A'
	}
	return r
}

func isVowel(r rune) bool {
	r = lower(r)
	switch r {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

func TranslateWord(word string) string {
	firstVowelIndex := -1

	word = strings.ToLower(word)

	for i, r := range word {
		if isVowel(r) {
			firstVowelIndex = i
			break
		}
	}
	switch {
	case firstVowelIndex == 0:
		return word + vowelSuffix
	case firstVowelIndex > 0:
		consonant := word[:firstVowelIndex]
		endOfWord := word[firstVowelIndex:]
		return endOfWord + consonant + consonantSuffix
	default:
		return word
	}
}

func TranslatePhrase(phrace string) string {
	words := strings.FieldsFunc(phrace, func(c rune) bool {
		return !unicode.IsLetter(c)

	})
	for _, w := range words {
		phrace = strings.Replace(phrace, w, TranslateWord(w), 1)
	}
	return phrace
}
