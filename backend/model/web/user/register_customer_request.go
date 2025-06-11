package user

import (
	"lampung_trip/helper"
	"lampung_trip/model/domain"
	"lampung_trip/model/web/customer"
)

type RegisterCustomerRequest struct {
	Username string                 `json:"username"`
	Password string                 `json:"password"`
	Customer customer.CreateRequest `json:"customer"`
}

func RegisterCustomerRequestToDomain(request RegisterCustomerRequest) domain.User {
	return domain.User{
		Username: request.Username,
		Password: helper.HashedPassword(request.Password),
		Role:     "Customer",
		Customer: customer.CreateRequestToDomain(request.Customer),
	}
}
