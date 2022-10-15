package sheets

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func NewClient(credentials []byte) *sheets.Service {
	ctx := context.Background()
	client, err := sheets.NewService(ctx, option.WithCredentialsJSON(credentials))
	if err != nil {
		fmt.Println(err)
	}

	return client
}
