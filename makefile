.PHONY: validate
validate:
	go run ./cmd/validate wordlist4096.txt

.PHONY: suggest
suggest:
	go run ./cmd/suggest

.PHONY: tidy
tidy:
	sort -u -o wordlist4096.txt wordlist4096.txt
