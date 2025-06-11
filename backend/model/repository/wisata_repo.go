package repository

import "lampung_trip/model/domain"

type WisataRepo interface {
	Create(wisata domain.Wisata) error
	Update(wisata domain.Wisata) error
	GetAll() ([]domain.Wisata, error)
	Delete(id int) error
}
