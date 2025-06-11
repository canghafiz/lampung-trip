package customer

import "lampung_trip/model/domain"

type CreateRequest struct {
	Nama  string `json:"nama"`
	Photo string `json:"photo"`
	NoHp  string `json:"no_hp"`
}

func CreateRequestToDomain(request CreateRequest) domain.Customer {
	return domain.Customer{
		Nama:  request.Nama,
		Photo: request.Photo,
		NoHp:  request.NoHp,
	}
}
