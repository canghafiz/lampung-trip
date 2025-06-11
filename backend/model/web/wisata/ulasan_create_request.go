package wisata

import "lampung_trip/model/domain"

type UlasanCreateRequest struct {
	WisataId   int     `json:"wisata_id"`
	CustomerId int     `json:"customer_id"`
	Rating     float64 `json:"rating"`
	Keterangan string  `json:"keterangan"`
}

func UlasanCreateRequestToDomain(request UlasanCreateRequest) domain.UlasanWisata {
	return domain.UlasanWisata{
		WisataId:   request.WisataId,
		CustomerId: request.CustomerId,
		Rating:     request.Rating,
		Keterangan: request.Keterangan,
	}
}
