package discord

import (
	"encoding/json"
	"log"
)

// Body is the body of the request received from Discord
type Body struct {
	Id     string `json:"id"`
	Token  string `json:"token"`
	Type   int    `json:"type"`
	Member struct {
		User `json:"user"`
	} `json:"member"`
	Data struct {
		CommandName string  `json:"name"`
		Entries     []Entry `json:"options"`
	}
}

// User is an object that is sent in the request, and contains information about who sent the request.
type User struct {
	Id   string `json:"id"`
	Name string `json:"username"`
}

// Entry contains information about the command options sent in the request.
type Entry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// ParseRequestBody unmarshalls the body of the request received from Discord, and returns a Body object.
func ParseRequestBody(body string) (Body, error) {
	var b Body

	err := json.Unmarshal([]byte(body), &b)
	if err != nil {
		return Body{}, err
	}
	log.Printf("Parsed Body: %+v\n", b)

	return b, nil
}
