package repository

import "lampung_trip/model/domain"

type UlasanWisataRepo interface {
	Create(ulasan domain.UlasanWisata) error
	Delete(ulasan domain.UlasanWisata) error
}
