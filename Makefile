include .env
export

APP_NAME=weather-monitor
CMD_DIR=./cmd/weather-monitor

.PHONY: run test

run:
	go run $(CMD_DIR)

test:
	go test ./...

