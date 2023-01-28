.PHONY: validate
validate:
	go test -v

.PHONY: suggest
suggest:
	go run ./cmd/suggest

.PHONY: tidy
tidy:
	sort -u -o wordlist4096.txt wordlist4096.txt
