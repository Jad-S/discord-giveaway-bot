package sheets

// Spreadsheet contains the spreadsheet ID and the range of the spreadsheet.
type Spreadsheet struct {
	Id    string `yaml:"id"`
	Range string `yaml:"range"`
}

// List contains the list of spreadsheets for the giveaways
type List struct {
	Guess  Spreadsheet `yaml:"guess-spreadsheet"`
	Raffle Spreadsheet `yaml:"raffle-spreadsheet"`
}
