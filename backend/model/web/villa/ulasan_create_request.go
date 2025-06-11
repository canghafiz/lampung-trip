package villa

import "lampung_trip/model/domain"

type UlasanCreateRequest struct {
	VillaId    int     `json:"villa_id"`
	CustomerId int     `json:"customer_id"`
	Rating     float64 `json:"rating"`
	Keterangan string  `json:"keterangan"`
}

func UlasanCreateRequestToDomain(request UlasanCreateRequest) domain.UlasanVilla {
	return domain.UlasanVilla{
		VillaId:    request.VillaId,
		CustomerId: request.CustomerId,
		Rating:     request.Rating,
		Keterangan: request.Keterangan,
	}
}
