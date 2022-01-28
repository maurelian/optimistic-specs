GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')
GITVERSION := $(shell echo 0.0.0)

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.GitVersion=$(GITVERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

bss:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/bss

clean:
	rm bss

test:
	go test -v ./...

lint:
	golangci-lint run ./...

.PHONY: \
	bss \
	clean \
	test \
	lint
