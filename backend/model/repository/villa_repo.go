package repository

import "lampung_trip/model/domain"

type VillaRepo interface {
	Create(villa domain.Villa) error
	Update(villa domain.Villa) error
	GetAll() ([]domain.Villa, error)
	Delete(id int) error
}
