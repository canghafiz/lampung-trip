package villa

import "lampung_trip/model/domain"

type UpdateRequest struct {
	Id        int                `json:"id"`
	Judul     string             `json:"judul"`
	Deskripsi string             `json:"deskripsi"`
	Gambar    []GambarRequest    `json:"gambar"`
	Info      InfoRequest        `json:"info"`
	Fasilitas []FasilitasRequest `json:"fasilitas"`
}

func UpdateRequestToDomain(request UpdateRequest) domain.Villa {
	return domain.Villa{
		Id:        request.Id,
		Judul:     request.Judul,
		Deskripsi: request.Deskripsi,
		Gambar:    GambarRequestToDomains(request.Gambar),
		Info:      InfoRequestToDomain(request.Info),
		Fasilitas: FasilitasRequestToDomains(request.Fasilitas),
	}
}
