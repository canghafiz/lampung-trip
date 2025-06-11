package service

import (
	"fmt"
	"lampung_trip/model/repository"
	"lampung_trip/model/web/villa"
)

type UlasanVillaServImpl struct {
	Repo repository.UlasanVillaRepo
}

func NewUlasanVillaServImpl(repo repository.UlasanVillaRepo) *UlasanVillaServImpl {
	return &UlasanVillaServImpl{Repo: repo}
}

func (serv *UlasanVillaServImpl) Create(request villa.UlasanCreateRequest) error {
	// Model
	model := villa.UlasanCreateRequestToDomain(request)

	// Call Repo
	err := serv.Repo.Create(model)
	if err != nil {
		return fmt.Errorf("create ulasan villa failed")
	}

	return nil
}

func (serv *UlasanVillaServImpl) Delete(request villa.UlasanDeleteRequest) error {
	// Model
	model := villa.UlasanDeleteRequestToDomain(request)

	// Call Repo
	err := serv.Repo.Delete(model)
	if err != nil {
		return fmt.Errorf("delete ulasan villa failed")
	}

	return nil
}
