package service

import "github.com/riceChuang/marryme/repo"

type MerryMeSvc struct {
	Storage *repo.GuestRepo
}

var instance *MerryMeSvc


// New new MerryMe svc
func NewMerryMeSvc(storage *repo.GuestRepo) *MerryMeSvc  {
	instance = &MerryMeSvc{
		Storage: storage,
	}
	return  instance
}

// Get MerryMe svc
func GetInstance() *MerryMeSvc {
	return instance
}

