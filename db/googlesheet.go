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

func NewGoogleSheet(config *config.Config) (service *GoogleSheet, err error) {

	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithAPIKey(config.APIKey))
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve Sheets client: %v", err)

	}

	return &GoogleSheet{Client: srv, SpreadSheetID: config.SpreadSheetID}, nil
}

func (gs *GoogleSheet) GetGuests() (guests []*types.Guest, err error) {

	readRange := "A2:G"
	resp, err := gs.Client.Spreadsheets.Values.Get(gs.SpreadSheetID, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve data from sheet: %v", err)
	}
	

	for _, row := range resp.Values {
		accompanies, err := strconv.ParseUint(row[2].(string), 10, 32)
		if err != nil {
			log.Infof("Ignore name: %s", row[0].(string))
			continue
		}

		money, err := strconv.ParseUint(row[3].(string), 10, 32)
		if err != nil {
			log.Infof("Ignore name: %s", row[0].(string))
			continue
		}

		isAttend := false
		if len(row) == 5 {
			tmpAttend := row[4].(string)
			if tmpAttend == "是" || tmpAttend == "y" || tmpAttend == "Y" {
				isAttend = true
			}
		}

		var updateTime, createTime time.Time

		if row[5] != nil && len(row[5].(string)) > 0 {
			updateTime, err = time.Parse("2006/01/02", row[5].(string))
			if err != nil {
				fmt.Println(err)
			}
		}

		if row[6] != nil && len(row[6].(string)) > 0 {
			createTime, err = time.Parse("2006/01/02", row[6].(string))
			if err != nil {
				fmt.Println(err)
			}
		}

		guests = append(guests, &types.Guest{
			Name:        row[0].(string),
			NickName:    row[1].(string),
			Accompanies: uint8(accompanies),
			Money:       uint(money),
			IsAttend:    isAttend,
			UpdateAt:    updateTime,
			CreatedAt:   createTime,
		})
	}
	return
}

func (gs *GoogleSheet) SaveGuest([]*types.Guest) (guests *types.Guest, err error) {
	gs.Client.Spreadsheets.BatchUpdate()
	
	data := []*sheets.ValueRange{}
	return 
}
