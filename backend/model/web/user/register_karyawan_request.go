package user

import (
	"lampung_trip/helper"
	"lampung_trip/model/domain"
	"lampung_trip/model/web/karyawan"
)

type RegisterKaryawanRequest struct {
	Username string                 `json:"username"`
	Password string                 `json:"password"`
	Karyawan karyawan.CreateRequest `json:"karyawan"`
}

func RegisterKaryawanRequestToDomain(request RegisterKaryawanRequest) domain.User {
	return domain.User{
		Username: request.Username,
		Password: helper.HashedPassword(request.Password),
		Role:     "Karyawan",
		Karyawan: karyawan.CreateRequestToDomain(request.Karyawan),
	}
}
