package piglatin

import (
	"strings"
	"unicode"
)

const (
	vowelSuffix     = "yay"
	consonantSuffix = "ay"
)

func translateWord(word string) string {
	if word == "" {
		return word
	}
	capitalazed := unicode.IsUpper(rune(word[0]))
	firstVowelIndex := strings.IndexFunc(word, isVowel)

	switch {
	case firstVowelIndex == 0:
		word += vowelSuffix
	case firstVowelIndex > 0:
		consonant := word[:firstVowelIndex]
		endOfWord := word[firstVowelIndex:]
		word = endOfWord + consonant + consonantSuffix
	}

	word = strings.ToLower(word)
	if capitalazed {
		word = strings.Title(word)
	}

	return word
}

// Translate translate english phrace or word into the piglatina.
func Translate(phrase string) string {
	words := strings.Fields(phrase)
	for i, w := range words {
		w = strings.TrimRightFunc(w, isNotLetter)
		words[i] = strings.Replace(words[i], w, translateWord(w), 1)
	}

	return strings.Join(words, " ")
}

func lower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
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

func isNotLetter(r rune) bool {
	return !unicode.IsLetter(r)
}
