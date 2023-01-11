package validate

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kklash/wordlist-4096/damlev"
)

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
