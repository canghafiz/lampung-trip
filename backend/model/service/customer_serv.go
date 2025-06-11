package service

import (
	"lampung_trip/model/web/customer"
)

type CustomerServ interface {
	UpdateData(request customer.UpdateDataRequest) error
	GetSingleByUserId(userId int) (*customer.Response, error)
	GetAll() ([]customer.Response, error)
}
