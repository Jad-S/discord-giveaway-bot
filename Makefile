build-DiscordGiveawayCommandsFunction:
	GOOS=linux GOARCH=amd64 go build -o handler cmd/main.go
	mv handler $(ARTIFACTS_DIR)/