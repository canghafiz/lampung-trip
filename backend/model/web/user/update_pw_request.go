package user

import (
	"lampung_trip/helper"
	"lampung_trip/model/domain"
)

type UpdatePwRequest struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func UpdatePwRequestToDomain(request UpdatePwRequest, id int) domain.User {
	return domain.User{
		Id:       id,
		Username: request.Username,
		Password: helper.HashedPassword(request.NewPassword),
	}
}
