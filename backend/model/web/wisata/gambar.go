package wisata

import "lampung_trip/model/domain"

type GambarResponse struct {
	Url string `json:"url"`
}

func ToGambarResponse(gambar domain.GambarWisata) *GambarResponse {
	return &GambarResponse{
		Url: gambar.Url,
	}
}

func ToGambarResponses(data []domain.GambarWisata) []GambarResponse {
	var gambars []GambarResponse

	for i := range data {
		gambars = append(gambars, *ToGambarResponse(data[i]))
	}

	return gambars
}

type GambarRequest struct {
	Url string `json:"url"`
}

func GambarRequestToDomain(request GambarRequest) domain.GambarWisata {
	return domain.GambarWisata{
		Url: request.Url,
	}
}

func GambarRequestToDomains(requests []GambarRequest) []domain.GambarWisata {
	var gambars []domain.GambarWisata

	for i := range requests {
		gambars = append(gambars, GambarRequestToDomain(requests[i]))
	}

	return gambars
}
