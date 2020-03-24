package repo

import (
	"github.com/riceChuang/marryme/config"
	"github.com/riceChuang/marryme/db"
	"github.com/riceChuang/marryme/types"
	"strings"
)

type GuestRepo struct {
	Guests []*types.Guest
}

func NewGuestsRepo(conf *config.Config) (repo *GuestRepo, err error) {

	googleSheet, err := db.NewGoogleSheet(conf)
	if err != nil {
		return nil, err
	}

	guests, err := googleSheet.GetGuests()
	if err != nil {
		return nil, err
	}

	repo = &GuestRepo{
		Guests: guests,
	}
	return
}

func (g *GuestRepo) SearchGuests(source string) []types.Guest {
	var matches []types.Guest

	for _, guest := range g.Guests {
		if match(source, guest.Name, guest.NickName) {
			matches = append(matches, *guest)
		}
	}
	return matches
}

func (g *GuestRepo) SaveGuests()(err error) {
	for _, guest := range g.Guests{
		
	}
}

func match(source string, name string, nickName string) bool {
	return strings.Contains(name, source) || strings.Contains(nickName, source)
}
