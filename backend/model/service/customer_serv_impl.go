package service

import (
	"fmt"
	"lampung_trip/model/repository"
	"lampung_trip/model/web/customer"
)

type CustomerServImpl struct {
	Repo repository.CustomerRepo
}

func NewCustomerServImpl(repo repository.CustomerRepo) *CustomerServImpl {
	return &CustomerServImpl{Repo: repo}
}

func (serv *CustomerServImpl) UpdateData(request customer.UpdateDataRequest) error {
	// Call repo
	domain := customer.UpdateDataRequestToDomain(request)
	err := serv.Repo.UpdateData(domain)
	if err != nil {
		return fmt.Errorf("failed to update data customer")
	}

	return nil
}

func (serv *CustomerServImpl) GetSingleByUserId(userId int) (*customer.Response, error) {
	// Call Repo
	customerRepo, err := serv.Repo.GetSingleByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get customer data")
	}

	return customer.ToResponse(*customerRepo, *customerRepo.User), nil
}

func (serv *CustomerServImpl) GetAll() ([]customer.Response, error) {
	// Call Repo
	customerRepo, err := serv.Repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get customer data")
	}

	return customer.ToResponses(customerRepo), nil
}
