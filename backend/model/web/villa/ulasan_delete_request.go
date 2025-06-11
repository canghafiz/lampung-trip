package villa

import "lampung_trip/model/domain"

type UlasanDeleteRequest struct {
	VillaId    int `json:"villa_id"`
	CustomerId int `json:"customer_id"`
}

func UlasanDeleteRequestToDomain(request UlasanDeleteRequest) domain.UlasanVilla {
	return domain.UlasanVilla{
		VillaId:    request.VillaId,
		CustomerId: request.CustomerId,
	}
}
