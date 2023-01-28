package wordlist4096

import (
	_ "embed"
	"strings"
)

// BitsPerWord is the number of bits of information represented by each word
// in the wordlist.
const BitsPerWord uint = 12

// WordList is the mnemonic encoding wordlist in alphabetical sorted order.
var WordList []string

// WordMap is a mapping of words to their indices in the wordlist.
var WordMap = make(map[string]uint16)

//go:embed wordlist4096.txt
var rawWordList string

func init() {
	WordList = strings.Fields(rawWordList)

	for i, word := range WordList {
		WordMap[word] = uint16(i)
	}
}
