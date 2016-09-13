test:
	go test -v ./... | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''

deps:
	go get -t -d -v

build: deps
	go build -v

clean:
	rm rabot

.DEFAULT_GOAL := build
