package user

import (
	"lampung_trip/helper"
	"lampung_trip/model/domain"
)

type RegisterAdminRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterAdminRequestToDomain(request RegisterAdminRequest) domain.User {
	return domain.User{
		Username: request.Username,
		Password: helper.HashedPassword(request.Password),
		Role:     "Admin",
	}
}
