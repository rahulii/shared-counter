.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: bench
bench:
	go test -bench=. ./...