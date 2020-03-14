package repo

import "time"

type Guest struct {
	Name        string
	NickName    string
	Accompanies uint8
	Money       uint
	IsAttend    bool
	CreatedAt   time.Time
}
