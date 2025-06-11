package customer

import "lampung_trip/model/domain"

type Response struct {
	Id       int    `json:"id"`
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Nama     string `json:"nama"`
	Photo    string `json:"photo"`
	NoHp     string `json:"no_hp"`
}

func ToResponse(customer domain.Customer, user domain.User) *Response {
	return &Response{
		Id:       customer.Id,
		UserId:   user.Id,
		Username: user.Username,
		Nama:     customer.Nama,
		Photo:    customer.Photo,
		NoHp:     customer.NoHp,
	}
}

func ToResponses(customers []domain.Customer) []Response {
	var responses []Response

	for _, customer := range customers {
		responses = append(responses, *ToResponse(customer, *customer.User))
	}

	return responses
}
