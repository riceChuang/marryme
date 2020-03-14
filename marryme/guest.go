package marryme

import (
	"log"
	"strconv"
	"strings"
	"time"
)

type Guest struct {
	Name        string
	NickName    string
	Accompanies uint8
	Money       uint
	IsAttend    bool
	CreatedAt   time.Time
}

var (
	guests []Guest
)

func InitialGuests(rows [][]interface{}) []Guest {
	guests = []Guest{}

	for _, row := range rows {
		accompanies, err := strconv.ParseUint(row[2].(string), 10, 32)
		if err != nil {
			log.Printf("Ignore name: %s", row[0].(string))
			continue
		}

		money, err := strconv.ParseUint(row[3].(string), 10, 32)
		if err != nil {
			log.Printf("Ignore name: %s", row[0].(string))
			continue
		}

		isAttend := false
		if len(row) == 5 {
			tmpAttend := row[4].(string)
			if tmpAttend == "是" || tmpAttend == "y" || tmpAttend == "Y" {
				isAttend = true
			}
		}

		guests = append(guests, Guest{
			Name:        row[0].(string),
			NickName:    row[1].(string),
			Accompanies: uint8(accompanies),
			Money:       uint(money),
			IsAttend:    isAttend,
		})
	}

	return guests
}

func GetGuests() []Guest {
	return guests
}

func SearchGuests(source string) []Guest {
	var matches []Guest

	for _, guest := range guests {
		if match(source, guest.Name, guest.NickName) {
			matches = append(matches, guest)
		}
	}

	return matches
}

func match(source string, name string, nickName string) bool {
	return strings.Contains(name, source) || strings.Contains(nickName, source)
}
