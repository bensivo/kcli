.PHONY: PHONY

build: .PHONY
	go build ./cmd/kcli

install: .PHONY
	go install ./cmd/kcli

test: .PHONY
	bats ./test
