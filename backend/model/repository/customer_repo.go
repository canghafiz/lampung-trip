package repository

import "lampung_trip/model/domain"

type CustomerRepo interface {
	UpdateData(customer domain.Customer) error
	GetSingleByUserId(userId int) (*domain.Customer, error)
	GetAll() ([]domain.Customer, error)
}
