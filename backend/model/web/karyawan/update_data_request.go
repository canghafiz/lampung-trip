package karyawan

import "lampung_trip/model/domain"

type UpdateDataRequest struct {
	Id    int    `json:"id"`
	Nama  string `json:"nama"`
	Photo string `json:"photo"`
	NoHp  string `json:"no_hp"`
}

func UpdateDataRequestToDomain(request UpdateDataRequest) domain.Karyawan {
	return domain.Karyawan{
		UserId: request.Id,
		Nama:   request.Nama,
		Photo:  request.Photo,
		NoHp:   request.NoHp,
	}
}
