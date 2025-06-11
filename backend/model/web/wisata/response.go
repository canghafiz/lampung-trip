package wisata

import "lampung_trip/model/domain"

type Response struct {
	Id        int                 `json:"id"`
	Judul     string              `json:"judul"`
	Deskripsi string              `json:"deskripsi"`
	Rating    float64             `json:"rating"`
	Gambar    []GambarResponse    `json:"gambar"`
	Fasilitas []FasilitasResponse `json:"fasilitas"`
	Info      InfoResponse        `json:"info"`
	Ulasan    []UlasanResponse    `json:"ulasan"`
}

func ToResponse(wisata domain.Wisata) *Response {
	return &Response{
		Id:        wisata.Id,
		Judul:     wisata.Judul,
		Deskripsi: wisata.Deskripsi,
		Rating:    wisata.Rating,
		Gambar:    ToGambarResponses(wisata.Gambar),
		Fasilitas: ToFasilitasResponses(wisata.Fasilitas),
		Info:      *ToInfoResponse(wisata.Info),
		Ulasan:    ToUlasanResponses(wisata.Ulasan),
	}
}

func ToResponses(wisata []domain.Wisata) []Response {
	var responses []Response

	for i := range wisata {
		responses = append(responses, *ToResponse(wisata[i]))
	}

	return responses
}
