package villa

import "lampung_trip/model/domain"

type InfoResponse struct {
	Harga  string `json:"harga"`
	Lokasi string `json:"lokasi"`
}

func ToInfoResponse(info domain.InfoVilla) *InfoResponse {
	return &InfoResponse{
		Harga:  info.Harga,
		Lokasi: info.Lokasi,
	}
}

type InfoRequest struct {
	WisataId int    `json:"wisata_id"`
	Harga    string `json:"harga"`
	Lokasi   string `json:"lokasi"`
}

func InfoRequestToDomain(request InfoRequest) domain.InfoVilla {
	return domain.InfoVilla{
		Id:     request.WisataId,
		Harga:  request.Harga,
		Lokasi: request.Lokasi,
	}
}
