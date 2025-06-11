package service

import (
	"fmt"
	"lampung_trip/model/repository"
	"lampung_trip/model/web/wisata"
)

type UlasanWisataServImpl struct {
	Repo repository.UlasanWisataRepo
}

func NewUlasanWisataServImpl(repo repository.UlasanWisataRepo) *UlasanWisataServImpl {
	return &UlasanWisataServImpl{Repo: repo}
}

func (serv *UlasanWisataServImpl) Create(request wisata.UlasanCreateRequest) error {
	// Model
	model := wisata.UlasanCreateRequestToDomain(request)

	// Call Repo
	err := serv.Repo.Create(model)
	if err != nil {
		return fmt.Errorf("create ulasan wisata failed")
	}

	return nil
}

func (serv *UlasanWisataServImpl) Delete(request wisata.UlasanDeleteRequest) error {
	// Model
	model := wisata.UlasanDeleteRequestToDomain(request)

	// Call Repo
	err := serv.Repo.Delete(model)
	if err != nil {
		return fmt.Errorf("create ulasan wisata failed")
	}

	return nil
}
