package wisata

import "lampung_trip/model/domain"

type InfoResponse struct {
	Harga      string `json:"harga"`
	WaktuBuka  string `json:"waktu_buka"`
	WaktuTutup string `json:"waktu_tutup"`
	HariBuka   string `json:"hari_buka"`
	HariTutup  string `json:"hari_tutup"`
	UrlLokasi  string `json:"url_lokasi"`
}

func ToInfoResponse(info domain.InfoWisata) *InfoResponse {
	return &InfoResponse{
		Harga:      info.Harga,
		WaktuBuka:  info.WaktuBuka,
		WaktuTutup: info.WaktuTutup,
		HariBuka:   info.HariBuka,
		HariTutup:  info.HariTutup,
		UrlLokasi:  info.UrlLokasi,
	}
}

type InfoRequest struct {
	Harga      string `json:"harga"`
	WaktuBuka  string `json:"waktu_buka"`
	WaktuTutup string `json:"waktu_tutup"`
	HariBuka   string `json:"hari_buka"`
	HariTutup  string `json:"hari_tutup"`
	UrlLokasi  string `json:"url_lokasi"`
}

func InfoRequestToDomain(request InfoRequest) domain.InfoWisata {
	return domain.InfoWisata{
		Harga:      request.Harga,
		WaktuBuka:  request.WaktuBuka,
		WaktuTutup: request.WaktuTutup,
		HariBuka:   request.HariBuka,
		HariTutup:  request.HariTutup,
		UrlLokasi:  request.UrlLokasi,
	}
}
