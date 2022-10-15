package discord

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

// VerifyDiscordRequest is called everytime a request is sent to validate the request is coming from Discord.
// Check this for more info: https://discord.com/developers/docs/interactions/slash-commands#security-and-authorization
func VerifyDiscordRequest(e events.APIGatewayV2HTTPRequest, publicKey string) (bool, error) {
	// Get the timestamp and signature from the headers, log values for debugging
	message := e.Body
	log.Printf("Message: %s", message)

	signature := e.Headers["x-signature-ed25519"]
	log.Printf("Signature: %s", signature)

	timestamp := e.Headers["x-signature-timestamp"]
	log.Printf("Timestamp: %s", timestamp)

	// Concatenate the timestamp and the message
	payload := timestamp + message
	fmt.Printf("Payload: %s", payload)

	// Convert the signature from hex to bytes
	sig, err := hex.DecodeString(signature)
	if err != nil {
		return false, fmt.Errorf("error decoding signature: %v", err)
	}

	// Convert the public key from hex to bytes
	ed25519Key, err := hex.DecodeString(publicKey)
	if err != nil {
		return false, fmt.Errorf("error decoding public key: %w", err)
	}

	// Verify that the request came from Discord
	return ed25519.Verify(ed25519Key, []byte(payload), sig), nil
}
