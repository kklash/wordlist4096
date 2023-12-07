# Wordlist 4096

This is a robust English wordlist for encoding data. It was inspired by previous natural language encoding systems, such as [What3Words](https://what3words.com/), [BIP-0039](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki), and [Diceware](https://diceware.dmuth.org/).

This wordlist was created for [the Mnemonikey project](https://github.com/kklash/mnemonikey).

## Specification

Each word in this list represents one of 4096 equally possible options. As per [Claude Shannon's theory of information](https://en.wikipedia.org/wiki/Information_theory), a word from this list encodes 12 bits of information, since $log_2(4096) = 12$ (i.e. $2^{12} = 4096$)

But this wordlist is not composed of any old random words. The list has been carefully constructed to fulfill the following properties:

|Property|Advantage|
|--------|---------|
|Avoids more than one homophone - words which sound the same but have different spelling. (e.g. `beat` and `beet`)|Avoids ambiguity and confusion when orally communicating words from the list.|
|Avoids more than one conceptually competitive word (e.g. `shook` and `shake` both refer to the same action in a different tense)|This prevents conceptual confusion, to allow words to be more easily memorized without mnemonic ambiguity.|
|Contains no words shorter than 3 characters or longer than 8 characters|Long words add complexity. Small words reduce uniqueness.|
|Every word is uniquely identifiable by at most the first four characters.|Allows fast automated interpretation when typing words as input.|
|The [Damerau-Levenshtein distance](https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance) between every word in the wordlist is at least 2.|Reduces the risk of typos which could convert one valid word into another valid one by accident.|

## Antecedents

This project was inspired by the following natural-language encoding projects:

- [What3Words](https://what3words.com/)
- [Diceware](https://diceware.dmuth.org/)
- [BIP39](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki)
- [SLIP39](https://github.com/satoshilabs/slips/blob/master/slip-0039.md)
- [PGP word list](https://en.wikipedia.org/wiki/PGP_word_list)
- [RFC 1751](https://www.rfc-editor.org/rfc/rfc1751)
- [RFC 1760](https://www.rfc-editor.org/rfc/rfc1760)
- [RFC 2289](https://www.rfc-editor.org/rfc/rfc2289)
- [mnemonicode](https://github.com/singpolyma/mnemonicode)

## Contributing

The tooling for this package requires a [Golang](https://go.dev) compiler.

For the moment, `wordlist4096` is not finalized. I'm happy to accept PRs to improve the wordlist. Please ensure your changes fulfill these requirements:

1. Any changes must pass tests. To run tests, use `go test` (or `make validate`).
2. Any new words must not contravene the heuristic properties discussed in the [Specification](#Specification) section.
3. Any word deletions or changes must be justifiable.

Be aware that discussions about what words are 'memorable' or 'confusing' may be highly subjective. In review, I may veto any arbitrary decision between words, simply to save time.

## Scripts

- To validate the computable properties of the wordlist, run `make validate`.
- To suggest new possible words to add to the wordlist, run `make suggest`. (This pulls from `/usr/share/dict/words`, only available on unix systems)
- To sort the wordlist and deduplicate words, run `make tidy`.
