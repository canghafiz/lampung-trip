package villa

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

func ToResponse(villa domain.Villa) *Response {
	return &Response{
		Id:        villa.Id,
		Judul:     villa.Judul,
		Deskripsi: villa.Deskripsi,
		Rating:    villa.Rating,
		Gambar:    ToGambarResponses(villa.Gambar),
		Fasilitas: ToFasilitasResponses(villa.Fasilitas),
		Info:      *ToInfoResponse(villa.Info),
		Ulasan:    ToUlasanResponses(villa.Ulasan),
	}
}

func ToResponses(villa []domain.Villa) []Response {
	var responses []Response

	for i := range villa {
		responses = append(responses, *ToResponse(villa[i]))
	}

	return responses
}
