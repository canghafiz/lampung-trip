package service

import (
	"fmt"
	"lampung_trip/model/repository"
	"lampung_trip/model/web/villa" // Assuming this package exists for web models
)

type VillaServImpl struct {
	Repo repository.VillaRepo
}

func NewVillaServImpl(repo repository.VillaRepo) *VillaServImpl {
	return &VillaServImpl{Repo: repo}
}

func (serv *VillaServImpl) Create(request villa.CreateRequest) error {
	// Model
	domain := villa.CreateRequestToDomain(request)

	// Call Repo
	err := serv.Repo.Create(domain)
	if err != nil {
		return fmt.Errorf("failed to create villa: %w", err)
	}

	return nil
}

func (serv *VillaServImpl) Update(request villa.UpdateRequest) error {
	// Model
	domain := villa.UpdateRequestToDomain(request)

	// Call Repo
	err := serv.Repo.Update(domain)
	if err != nil {
		return fmt.Errorf("failed to update villa: %w", err)
	}

	return nil
}

func (serv *VillaServImpl) GetAll() ([]villa.Response, error) {
	// Call Repo
	results, err := serv.Repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all villas: %w", err)
	}

	return villa.ToResponses(results), nil
}

func (serv *VillaServImpl) Delete(id int) error {
	// Call Repo
	err := serv.Repo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete villa: %w", err)
	}

	return nil
}
