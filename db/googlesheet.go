package db

import (
	"context"
	"fmt"
	"github.com/riceChuang/marryme/config"
	"github.com/riceChuang/marryme/types"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"strconv"
	"time"
)

type GoogleSheet struct {
	Client        *sheets.Service
	SpreadSheetID string
}

var googleSheet *GoogleSheet

func NewGoogleSheet(config *config.Config) (service *GoogleSheet, err error) {

	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithAPIKey(config.APIKey))
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve Sheets client: %v", err)

	}

	return &GoogleSheet{Client: srv, SpreadSheetID: config.SpreadSheetID}, nil
}

func (gs *GoogleSheet) GetGuests() (guests map[int]*types.Guest, err error) {
	readRange := "A2:G"
	resp, err := gs.Client.Spreadsheets.Values.Get(gs.SpreadSheetID, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve data from sheet: %v", err)
	}
	return gs.SheetToGuest(resp.Values)
}

func (gs *GoogleSheet) SheetToGuest(values [][]interface{}) (guests map[int]*types.Guest, err error) {
	guests = map[int]*types.Guest{}
	for _, row := range values {
		id, err := strconv.ParseUint(row[0].(string), 10, 32)
		if err != nil {
			log.Infof("Ignore name: %s", row[0].(string))
			continue
		}

		accompanies, err := strconv.ParseUint(row[3].(string), 10, 32)
		if err != nil {
			log.Infof("Ignore name: %s", row[0].(string))
			continue
		}

		money, err := strconv.ParseUint(row[4].(string), 10, 32)
		if err != nil {
			log.Infof("Ignore name: %s", row[0].(string))
			continue
		}

		var updateTime, createTime time.Time

		if row[6] != nil && len(row[6].(string)) > 0 {
			updateTime, err = time.Parse("2006/01/02", row[5].(string))
			if err != nil {
				fmt.Println(err)
			}
		}

		if row[7] != nil && len(row[7].(string)) > 0 {
			createTime, err = time.Parse("2006/01/02", row[6].(string))
			if err != nil {
				fmt.Println(err)
			}
		}

		guests[int(id)] = &types.Guest{
			ID:          int(id),
			Name:        row[1].(string),
			NickName:    row[2].(string),
			Accompanies: int(accompanies),
			Money:       int(money),
			IsAttend:    row[5].(string),
			UpdateAt:    updateTime,
			CreatedAt:   createTime,
		}
	}
	return
}

func (gs *GoogleSheet) GuestToSheet(guests []*types.Guest) (values [][]interface{}, err error) {

	for _, v := range guests {
		var guestValue = []interface{}{"", "", "", "", "N", "", ""}
		if len(v.Name) > 0 {
			guestValue[0] = v.Name
		}
		if len(v.NickName) > 0 {
			guestValue[1] = v.NickName
		}
		if v.Accompanies > 0 {
			guestValue[2] = string(v.Accompanies)
		}
		if v.Money > 0 {
			guestValue[3] = string(v.Money)
		}
		if len(v.IsAttend) > 0 {
			guestValue[4] = v.IsAttend
		}
		if !v.UpdateAt.IsZero() {
			guestValue[5] = v.UpdateAt.Format("2006/01/02")
		}

		if !v.CreatedAt.IsZero() {
			guestValue[5] = v.CreatedAt.Format("2006/01/02")
		}
		values = append(values, guestValue)
	}
	return
}

func (gs *GoogleSheet) SaveGuest(guests []*types.Guest) (err error) {
	ctx := context.Background()
	rangeData := "A2:G"
	updateValues, err := gs.GuestToSheet(guests)
	if err != nil {
		return
	}

	rb := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "USER_ENTERED",
	}
	rb.Data = append(rb.Data, &sheets.ValueRange{
		Range:  rangeData,
		Values: updateValues,
	})

	_, err = gs.Client.Spreadsheets.Values.BatchUpdate(gs.SpreadSheetID, rb).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	return
}
