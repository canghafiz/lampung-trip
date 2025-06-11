package repository

import "lampung_trip/model/domain"

type UlasanVillaRepo interface {
	Create(ulasan domain.UlasanVilla) error
	Delete(ulasan domain.UlasanVilla) error
}
