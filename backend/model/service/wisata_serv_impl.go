package service

import (
	"fmt"
	"lampung_trip/model/repository"
	"lampung_trip/model/web/wisata"
)

type WisataServImpl struct {
	Repo repository.WisataRepo
}

func NewWisataServImpl(repo repository.WisataRepo) *WisataServImpl {
	return &WisataServImpl{Repo: repo}
}

func (serv *WisataServImpl) Create(request wisata.CreateRequest) error {
	// Call Repo
	domain := wisata.CreateRequestToDomain(request)
	err := serv.Repo.Create(domain)
	if err != nil {
		return fmt.Errorf("failed to create wisata")
	}

	return nil
}

func (serv *WisataServImpl) Update(request wisata.UpdateRequest) error {
	// Call Repo
	domain := wisata.UpdateRequestToDomain(request)
	err := serv.Repo.Update(domain)
	if err != nil {
		return fmt.Errorf("failed to update wisata")
	}

	return nil
}

func (serv *WisataServImpl) GetAll() ([]wisata.Response, error) {
	// Call Repo
	results, err := serv.Repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get wisata")
	}

	return wisata.ToResponses(results), nil
}

func (serv *WisataServImpl) Delete(id int) error {
	// Call Repo
	err := serv.Repo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete wisata")
	}

	return nil
}
