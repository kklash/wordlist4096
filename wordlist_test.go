package wordlist4096

import (
	"sort"
	"testing"

	"github.com/kklash/wordlist4096/validate"
)

func TestSorting(t *testing.T) {
	sortedWordList := append([]string{}, WordList...)
	sort.Strings(sortedWordList)
	for i, word := range WordList {
		if sortedWordList[i] != word {
			t.Errorf("wordlist file is not sorted; run 'make sort'")
			return
		}
	}
}

func TestWordListLength(t *testing.T) {
	if len(WordList) != 4096 {
		t.Errorf("wordlist length is %d, not 4096", len(WordList))
	}
}

func TestWordLengths(t *testing.T) {
	for _, word := range WordList {
		if err := validate.WordLength(word); err != nil {
			t.Errorf("%s: %q", err, word)
		}
	}
}

func TestPrefixes(t *testing.T) {
	checkedWords := make([]string, 0, len(WordList))

	for _, word := range WordList {
		if err := validate.WordPrefix(word, checkedWords); err != nil {
			t.Error(err)
		}
		checkedWords = append(checkedWords, word)
	}
}

func TestDistances(t *testing.T) {
	checkedWords := make([]string, 0, len(WordList))

	for _, word := range WordList {
		if err := validate.WordMinDistance(word, checkedWords); err != nil {
			t.Error(err)
		}
		checkedWords = append(checkedWords, word)
	}
}

func TestDuplicates(t *testing.T) {
	if err := validate.DuplicateWords(WordList); err != nil {
		t.Error(err)
	}
}
