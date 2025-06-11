package villa

import "lampung_trip/model/domain"

type FasilitasResponse struct {
	UrlIcon string `json:"url_icon"`
	Judul   string `json:"judul"`
}

func ToFasilitasResponse(fasilitas domain.FasilitasVilla) *FasilitasResponse {
	return &FasilitasResponse{
		UrlIcon: fasilitas.UrlIcon,
		Judul:   fasilitas.Judul,
	}
}

func ToFasilitasResponses(data []domain.FasilitasVilla) []FasilitasResponse {
	var fasilitases []FasilitasResponse

	for i := range data {
		fasilitases = append(fasilitases, *ToFasilitasResponse(data[i]))
	}

	return fasilitases
}

type FasilitasRequest struct {
	UrlIcon string `json:"url_icon"`
	Judul   string `json:"judul"`
}

func FasilitasRequestToDomain(request FasilitasRequest) domain.FasilitasVilla {
	return domain.FasilitasVilla{
		UrlIcon: request.UrlIcon,
		Judul:   request.Judul,
	}
}

func FasilitasRequestToDomains(requests []FasilitasRequest) []domain.FasilitasVilla {
	var fasilitases []domain.FasilitasVilla

	for i := range requests {
		fasilitases = append(fasilitases, FasilitasRequestToDomain(requests[i]))
	}

	return fasilitases
}
