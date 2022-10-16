build-DiscordGiveawayCommandsFunction:
	GOOS=linux GOARCH=amd64 go build -o handleRequest cmd/handleRequest/main.go
	mv handleRequest $(ARTIFACTS_DIR)/