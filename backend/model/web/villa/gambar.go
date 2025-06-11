package villa

import "lampung_trip/model/domain"

type GambarResponse struct {
	Url string `json:"url"`
}

func ToGambarResponse(gambar domain.GambarVilla) *GambarResponse {
	return &GambarResponse{
		Url: gambar.Url,
	}
}

func ToGambarResponses(data []domain.GambarVilla) []GambarResponse {
	var gambars []GambarResponse

	for i := range data {
		gambars = append(gambars, *ToGambarResponse(data[i]))
	}

	return gambars
}

type GambarRequest struct {
	VillaId int    `json:"villa_id"`
	Url     string `json:"url"`
}

func GambarRequestToDomain(request GambarRequest) domain.GambarVilla {
	return domain.GambarVilla{
		VillaId: request.VillaId,
		Url:     request.Url,
	}
}

func GambarRequestToDomains(requests []GambarRequest) []domain.GambarVilla {
	var gambars []domain.GambarVilla

	for i := range requests {
		gambars = append(gambars, GambarRequestToDomain(requests[i]))
	}

	return gambars
}
