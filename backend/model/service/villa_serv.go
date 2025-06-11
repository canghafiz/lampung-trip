package service

import (
	"lampung_trip/model/web/villa"
)

type VillaServ interface {
	Create(request villa.CreateRequest) error
	Update(request villa.UpdateRequest) error
	GetAll() ([]villa.Response, error)
	Delete(id int) error
}
