// Package validate contains wordlist validation rules.
package validate

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kklash/wordlist4096/damlev"
)

// DuplicateWords returns an error if the given word list contains any duplicates.
func DuplicateWords(words []string) error {
	allWordsSet := make(map[string]struct{})
	for _, word := range words {
		if _, ok := allWordsSet[word]; ok {
			return fmt.Errorf("found duplicate word: %q", word)
		}
		allWordsSet[word] = struct{}{}
	}
	return nil
}

// WordLength returns an error if the given word is shorter than 3 characters
// or longer than 8 characters.
func WordLength(word string) error {
	l := len(word)
	if l < 3 {
		return errors.New("word is too short")
	}
	if l > 8 {
		return errors.New("word is too long")
	}
	return nil
}

// WordPrefix returns an error if the given word shares the same 4-character prefix
// as another word in the word list.
func WordPrefix(word string, allWords []string) error {
	prefix := word
	if len(word) >= 4 {
		prefix = word[:4]
	}

	for _, otherWord := range allWords {
		if strings.HasPrefix(otherWord, prefix) {
			return fmt.Errorf("found repeated prefix %q: %q and %q", prefix, word, otherWord)
		}
	}
	return nil
}

// WordMinDistance returns an error if the given word has a Damerau-Levenshtein distance of
// less than 2 from any other word in the word list.
func WordMinDistance(word string, words []string) error {
	minDistance := len(word)
	for _, otherWord := range words {
		distance := damlev.DamerauLevenshteinDistance(word, otherWord)
		if distance < minDistance {
			minDistance = distance
			if minDistance < 2 {
				return fmt.Errorf("words within distance 2: %q and %q", word, otherWord)
			}
		}
	}
	return nil
}

// All runs all validation checks on a word list.
func All(word string, currentWords []string) error {
	if err := WordLength(word); err != nil {
		return err
	}
	if err := WordPrefix(word, currentWords); err != nil {
		return err
	}
	if err := WordMinDistance(word, currentWords); err != nil {
		return err
	}
	return nil
}
