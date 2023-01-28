package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/kklash/wordlist4096/validate"
)

var plainWordRegExp = regexp.MustCompile("^[a-z]*$")

func readWordList(fpath string) ([]string, error) {
	fileData, err := os.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	allWords := strings.Fields(strings.TrimSpace(string(fileData)))
	return allWords, nil
}

func run() error {
	allEnglishWords, err := readWordList("/usr/share/dict/words")
	if err != nil {
		return err
	}
	allWords, err := readWordList("wordlist4096.txt")
	if err != nil {
		return err
	}

	for _, word := range allEnglishWords {
		word = strings.ToLower(word)
		if !plainWordRegExp.MatchString(word) {
			continue
		}
		if err := validate.All(word, allWords); err == nil {
			fmt.Println(word)
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}
