package types

import "time"

type Guest struct {
	Name        string
	NickName    string
	Accompanies uint8
	Money       uint
	IsAttend    bool
	UpdateAt    time.Time
	CreatedAt   time.Time
}
