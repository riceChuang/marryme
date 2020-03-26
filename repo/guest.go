package repo

import (
	"github.com/riceChuang/marryme/config"
	"github.com/riceChuang/marryme/db"
	"github.com/riceChuang/marryme/types"
	"time"
)

type GuestRepo struct {
	Guests      map[int]*types.Guest
	GoogleSheet *db.GoogleSheet
}

func NewGuestsRepo(googleSheet *db.GoogleSheet) (repo *GuestRepo, err error) {
	guests, err := googleSheet.GetGuests()
	if err != nil {
		return nil, err
	}
	repo = &GuestRepo{
		Guests:      guests,
		GoogleSheet: googleSheet,
	}
	return
}

func (g *GuestRepo) SearchGuests(ids []int) (result []*types.Guest) {
	result = []*types.Guest{}
	for _, id := range ids {
		if guest, ok := g.Guests[id]; ok {
			result = append(result, guest)
		}
	}
	return result
}

func (g *GuestRepo) CreateOrUpdateGuest(guest *types.Guest) (err error) {
	if v, ok := g.Guests[guest.ID]; ok {
		if v.Name != guest.Name {
			v.Name = guest.Name
		}
		if v.NickName != guest.NickName {
			v.NickName = guest.NickName
		}
		if v.Money != guest.Money {
			v.Money = guest.Money
		}
		if v.Accompanies != guest.Accompanies {
			v.Accompanies = guest.Accompanies
		}
		if guest.IsAttend != v.IsAttend {
			v.IsAttend = guest.IsAttend
		}
		v.UpdateAt = time.Now()
	} else {
		guest.ID = len(g.Guests) + 1
		g.Guests[guest.ID] = guest
	}
	return
}

func (g *GuestRepo) SyncGuestToGoogleSheet() (err error) {

}
