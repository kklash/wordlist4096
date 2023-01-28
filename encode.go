package wordlist4096

import (
	"errors"
	"math/big"
	"math/bits"
)

// BitsPerWord is the number of bits of information represented by each word
// in the wordlist.
const BitsPerWord uint = 12

var wordMask = big.NewInt(int64((1 << BitsPerWord) - 1))

// ErrInvalidIndex is returned when given a word index which
// cannot be represented in at most BitsPerWord bits.
var ErrInvalidIndex = errors.New("word index value is too large")

// EncodeToIndices encodes the given payload as a slice of word indices.
// The bitSize parameter tells
func EncodeToIndices(payloadInt *big.Int, bitSize uint) []uint16 {
	nWords := (bitSize + BitsPerWord - 1) / BitsPerWord
	indices := make([]uint16, nWords)
	for i := nWords - 1; payloadInt.BitLen() > 0; i-- {
		indices[i] = uint16(new(big.Int).And(payloadInt, wordMask).Uint64())
		payloadInt.Rsh(payloadInt, BitsPerWord)
	}
	return indices
}

// EncodeToWords encodes the given word indices as a series of words from
// the wordlist. The returned phrase is encoded with lower-case characters.
//
// Returns ErrInvalidIndex if any of the given indices are outside the
// domain of the word list.
func EncodeToWords(indices []uint16) ([]string, error) {
	words := make([]string, len(indices))
	for i, index := range indices {
		if uint(bits.Len16(index)) > BitsPerWord {
			return nil, ErrInvalidIndex
		}
		words[i] = WordList[index]
	}
	return words, nil
}
