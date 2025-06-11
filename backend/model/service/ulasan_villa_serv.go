package service

import (
	"lampung_trip/model/web/villa"
)

type UlasanVillaServ interface {
	Create(request villa.UlasanCreateRequest) error
	Delete(request villa.UlasanDeleteRequest) error
}
