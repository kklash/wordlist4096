package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/kklash/wordlist4096/validate"
)

func CheckWordLengths(words []string) (ok bool) {
	ok = true
	fmt.Print("checking word lengths... ")
	defer func() {
		if ok {
			fmt.Println("DONE")
		} else {
			fmt.Println("\nDONE")
		}
	}()
	for _, word := range words {
		if err := validate.WordLength(word); err != nil {
			ok = false
			fmt.Printf("\n%s: %q", err, word)
		}
	}
	return
}

func CheckPrefixes(words []string) (ok bool) {
	ok = true
	fmt.Print("checking unique 4-char prefixes... ")
	defer func() {
		if ok {
			fmt.Println("DONE")
		} else {
			fmt.Println("\nDONE")
		}
	}()

	checkedWords := make([]string, 0, len(words))

	for _, word := range words {
		if err := validate.WordPrefix(word, checkedWords); err != nil {
			ok = false
			fmt.Printf("\n%s", err)
		}
		checkedWords = append(checkedWords, word)
	}
	return
}

func CheckDistances(words []string) (ok bool) {
	ok = true
	fmt.Print("checking Damerau-Levenshtein distances... ")
	defer func() {
		if ok {
			fmt.Println("DONE")
		} else {
			fmt.Println("\nDONE")
		}
	}()
	checkedWords := make([]string, 0, len(words))

	for _, word := range words {
		if err := validate.WordMinDistance(word, checkedWords); err != nil {
			ok = false
			fmt.Printf("\n%s", err)
		}
		checkedWords = append(checkedWords, word)
	}
	return
}

func run() error {
	if len(os.Args) < 2 {
		return errors.New("must provide wordlist file path")
	}
	wordFileData, err := os.ReadFile(os.Args[1])
	if err != nil {
		return err
	}

	lines := strings.Split(strings.TrimSpace(string(wordFileData)), "\n")
	words := make([]string, 0, len(lines))
	for _, line := range lines {
		word := strings.TrimSpace(line)
		if word != "" {
			words = append(words, word)
		}
	}

	allOK := true

	if err := validate.DuplicateWords(words); err != nil {
		allOK = false
		fmt.Println(err)
	}

	if ok := CheckWordLengths(words); !ok {
		allOK = false
	}
	if ok := CheckPrefixes(words); !ok {
		allOK = false
	}
	if ok := CheckDistances(words); !ok {
		allOK = false
	}

	if !allOK {
		return errors.New("some checks failed")
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}
