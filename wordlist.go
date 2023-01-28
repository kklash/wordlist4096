package wordlist4096

import (
	_ "embed"
	"strings"
)

var (
	// WordList is the mnemonic encoding wordlist in alphabetical sorted order.
	WordList []string

	// WordMap is a mapping of words to their indices in the wordlist.
	WordMap = make(map[string]uint16)
)

//go:embed wordlist4096.txt
var rawWordList string

func init() {
	WordList = strings.Fields(rawWordList)

	for i, word := range WordList {
		WordMap[word] = uint16(i)
	}
}
