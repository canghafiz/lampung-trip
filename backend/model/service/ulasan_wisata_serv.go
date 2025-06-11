package service

import "lampung_trip/model/web/wisata"

type UlasanWisataServ interface {
	Create(request wisata.UlasanCreateRequest) error
	Delete(request wisata.UlasanDeleteRequest) error
}
