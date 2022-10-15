package sheets

import (
	"fmt"
	"giveaway/internal/discord"
	"google.golang.org/api/sheets/v4"
	"log"
)

// VerifyEntry checks if the user has already entered the raffle by verifying whether the ID exists in the spreadsheet.
func VerifyEntry(client *sheets.Service, sheet Spreadsheet, id string) bool {
	exists, _ := client.Spreadsheets.Values.Get(sheet.Id, sheet.Range).Do()
	for _, row := range exists.Values {
		if row[0] == id {
			return true
		}
	}

	return false
}

func AddNewEntry(client *sheets.Service, body discord.Body, sheet Spreadsheet) error {
	switch body.Data.CommandName {
	case "raffle":
		err := NewRaffleEntry(client, sheet, body)
		if err != nil {
			return err
		}
	case "guess":
		err := NewGuessEntry(client, sheet, body)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewRaffleEntry(client *sheets.Service, sheet Spreadsheet, body discord.Body) error {
	discordId := body.Member.User.Id
	discordName := body.Member.User.Name

	twitchName := body.Data.Entries[0].Value
	casinoName := body.Data.Entries[1].Value

	// new entry details
	entry := &sheets.ValueRange{
		Values: [][]interface{}{
			{discordId, twitchName, casinoName, discordName},
		},
	}

	// add the new entry to the spreadsheet
	resp, err := client.Spreadsheets.Values.Append(sheet.Id, sheet.Range, entry).Do()
	if err != nil {
		return fmt.Errorf("unable to retrieve data from sheet: %v", err)
	}
	fmt.Printf("Updated cells: %s with values: %s", resp.Updates.UpdatedRange, entry.Values)

	return nil
}

func NewGuessEntry(client *sheets.Service, sheet Spreadsheet, body discord.Body) error {
	discordId := body.Member.User.Id
	discordName := body.Member.User.Name

	twitchName := body.Data.Entries[0].Value
	guessNumber := body.Data.Entries[1].Value

	// new entry details
	entry := &sheets.ValueRange{
		Values: [][]interface{}{
			{discordId, twitchName, guessNumber, discordName},
		},
	}

	// add the new entry to the spreadsheet
	resp, err := client.Spreadsheets.Values.Append(sheet.Id, sheet.Range, entry).Do()
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Printf("Updated cells: %s with values: %s", resp.Updates.UpdatedRange, entry.Values)

	return nil
}
