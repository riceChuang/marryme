package types

import "time"

type Guest struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	NickName    string    `json:"nick_name"`
	Accompanies int       `json:"accompanies"`
	Money       int       `json:"money"`
	IsAttend    string    `json:"is_attend"`
	UpdateAt    time.Time `json:"update_at"`
	CreatedAt   time.Time `json:"created_at"`
}
