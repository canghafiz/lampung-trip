package wisata

import "lampung_trip/model/domain"

type UlasanDeleteRequest struct {
	WisataId   int `json:"wisata_id"`
	CustomerId int `json:"customer_id"`
}

func UlasanDeleteRequestToDomain(request UlasanDeleteRequest) domain.UlasanWisata {
	return domain.UlasanWisata{
		WisataId:   request.WisataId,
		CustomerId: request.CustomerId,
	}
}
