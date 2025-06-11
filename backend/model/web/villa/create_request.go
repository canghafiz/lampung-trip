package villa

import "lampung_trip/model/domain"

type CreateRequest struct {
	Judul     string             `json:"judul"`
	Deskripsi string             `json:"deskripsi"`
	Gambar    []GambarRequest    `json:"gambar"`
	Info      InfoRequest        `json:"info"`
	Fasilitas []FasilitasRequest `json:"fasilitas"`
}

func CreateRequestToDomain(request CreateRequest) domain.Villa {
	return domain.Villa{
		Judul:     request.Judul,
		Deskripsi: request.Deskripsi,
		Rating:    0,
		Gambar:    GambarRequestToDomains(request.Gambar),
		Info:      InfoRequestToDomain(request.Info),
		Fasilitas: FasilitasRequestToDomains(request.Fasilitas),
	}
}
