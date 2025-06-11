package karyawan

import "lampung_trip/model/domain"

type Response struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Nama     string `json:"nama"`
	Photo    string `json:"photo"`
	NoHp     string `json:"no_hp"`
}

func ToResponse(karyawan domain.Karyawan, user domain.User) *Response {
	return &Response{
		UserId:   karyawan.UserId,
		Username: user.Username,
		Nama:     karyawan.Nama,
		Photo:    karyawan.Photo,
		NoHp:     karyawan.NoHp,
	}
}

func ToResponses(karyawans []domain.Karyawan) []Response {
	var responses []Response

	for _, karyawan := range karyawans {
		responses = append(responses, *ToResponse(karyawan, *karyawan.User))
	}

	return responses
}
