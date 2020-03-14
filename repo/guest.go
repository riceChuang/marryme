package repo

import (
	"fmt"
	"github.com/riceChuang/marryme/config"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"

	"context"
)

type GuestRepo struct {
	Guests []*Guest
}

func NewGuestsRepo(conf *config.Config) ( repo *GuestRepo, err error) {

	guests := []*Guest{}
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithAPIKey(conf.APIKey))
	if err != nil {
		return nil,fmt.Errorf("Unable to retrieve Sheets client: %v", err)
	}

	readRange := "Class Data!A2:E"
	resp, err := srv.Spreadsheets.Values.Get(conf.SpreadSheetID, readRange).Do()
	if err != nil {
		return nil,fmt.Errorf("Unable to retrieve data from sheet: %v", err)
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

		guests = append(guests, &Guest{
			Name:        row[0].(string),
			NickName:    row[1].(string),
			Accompanies: uint8(accompanies),
			Money:       uint(money),
			IsAttend:    isAttend,
		})
	}

	repo = &GuestRepo{
		Guests: guests,
	}
	return
}

func (g *GuestRepo) SearchGuests(source string) []Guest {
	var matches []Guest

	for _, guest := range g.Guests {
		if match(source, guest.Name, guest.NickName) {
			matches = append(matches, *guest)
		}
	}
	return matches
}

func match(source string, name string, nickName string) bool {
	return strings.Contains(name, source) || strings.Contains(nickName, source)
}
