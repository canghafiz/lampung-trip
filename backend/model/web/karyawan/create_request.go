package karyawan

import "lampung_trip/model/domain"

type CreateRequest struct {
	Nama  string `json:"nama"`
	Photo string `json:"photo"`
	NoHp  string `json:"no_hp"`
}

func CreateRequestToDomain(request CreateRequest) domain.Karyawan {
	return domain.Karyawan{
		Nama:  request.Nama,
		Photo: request.Photo,
		NoHp:  request.NoHp,
	}
}
