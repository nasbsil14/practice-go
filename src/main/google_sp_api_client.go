package main

import (
	"google.golang.org/api/sheets/v4"
	"fmt"
	"net/http"
	"golang.org/x/net/context"
	_"golang.org/x/oauth2"
	_"golang.org/x/oauth2/google"
	"log"
	"errors"
)

func main() {
	ctx := context.Background()

	c, err := getClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	sheetsService, err := sheets.New(c)
	if err != nil {
		log.Fatal(err)
	}

	// The ID of the spreadsheet to retrieve data from.
	spreadsheetId := "1CdZQCMvP3WbWpx5gJI6AflxneMGBsa9Ah4UxqdLRGAk" // TODO: Update placeholder value.

	// The A1 notation of the values to retrieve.
	range2 := "" // TODO: Update placeholder value.

	// How values should be represented in the output.
	// The default render option is ValueRenderOption.FORMATTED_VALUE.
	valueRenderOption := "" // TODO: Update placeholder value.

	// How dates, times, and durations should be represented in the output.
	// This is ignored if value_render_option is
	// FORMATTED_VALUE.
	// The default dateTime render option is [DateTimeRenderOption.SERIAL_NUMBER].
	dateTimeRenderOption := "" // TODO: Update placeholder value.

	resp, err := sheetsService.Spreadsheets.Values.Get(spreadsheetId, range2).ValueRenderOption(valueRenderOption).DateTimeRenderOption(dateTimeRenderOption).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Change code below to process the `resp` object:
	fmt.Printf("%#v\n", resp)
}

func getClient(ctx context.Context) (*http.Client, error) {
	// TODO: Change placeholder below to get authentication credentials. See
	// https://developers.google.com/sheets/quickstart/go#step_3_set_up_the_sample
	//
	// Authorize using the following scopes:
	//     sheets.DriveScope
	//     sheets.DriveReadonlyScope
	//     sheets.SpreadsheetsScope
	//     sheets.SpreadsheetsReadonlyScope
	return nil, errors.New("not implemented")
}
