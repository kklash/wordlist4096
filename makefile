.PHONY: validate
validate:
	go run ./cmd/validate wordlist-4096.txt

.PHONY: suggest
suggest:
	go run ./cmd/suggest

.PHONY: tidy
tidy:
	sort -u -o wordlist-4096.txt wordlist-4096.txt
