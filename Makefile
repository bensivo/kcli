.PHONY: PHONY

build: .PHONY
	go build 

install: .PHONY
	go install

test: .PHONY
	bats ./test
