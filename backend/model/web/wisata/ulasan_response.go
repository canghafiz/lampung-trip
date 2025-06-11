package wisata

import (
	"lampung_trip/model/domain"
	"lampung_trip/model/web/customer"
	"time"
)

type UlasanResponse struct {
	Customer   customer.Response `json:"customer"`
	Rating     float64           `json:"rating"`
	Keterangan string            `json:"keterangan"`
	CreatedAt  time.Time         `json:"created_at"`
}

func ToUlasanResponse(ulasan domain.UlasanWisata) *UlasanResponse {
	return &UlasanResponse{
		Customer:   *customer.ToResponse(ulasan.Customer, *ulasan.Customer.User),
		Rating:     ulasan.Rating,
		Keterangan: ulasan.Keterangan,
		CreatedAt:  ulasan.CreatedAt,
	}
}

func ToUlasanResponses(data []domain.UlasanWisata) []UlasanResponse {
	var ulasans []UlasanResponse

	for i := range data {
		ulasans = append(ulasans, *ToUlasanResponse(data[i]))
	}

	return ulasans
}
