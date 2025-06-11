package service

import (
	"lampung_trip/model/web/wisata"
)

type WisataServ interface {
	Create(request wisata.CreateRequest) error
	Update(request wisata.UpdateRequest) error
	GetAll() ([]wisata.Response, error)
	Delete(id int) error
}
