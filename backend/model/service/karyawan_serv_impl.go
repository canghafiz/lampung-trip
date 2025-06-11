package service

import (
	"fmt"
	"lampung_trip/model/repository"
	"lampung_trip/model/web/karyawan"
)

type KaryawanServImpl struct {
	Repo repository.KaryawanRepo
}

func NewKaryawanServImpl(repo repository.KaryawanRepo) *KaryawanServImpl {
	return &KaryawanServImpl{Repo: repo}
}

func (serv *KaryawanServImpl) UpdateData(request karyawan.UpdateDataRequest) error {
	// Call repo
	domain := karyawan.UpdateDataRequestToDomain(request)
	err := serv.Repo.UpdateData(domain)
	if err != nil {
		return fmt.Errorf("failed to update data karyawan")
	}

	return nil
}

func (serv *KaryawanServImpl) GetSingleByUserId(userId int) (*karyawan.Response, error) {
	// Call Repo
	karyawanRepo, err := serv.Repo.GetSingleByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get karyawan data")
	}

	return karyawan.ToResponse(*karyawanRepo, *karyawanRepo.User), nil
}

func (serv *KaryawanServImpl) GetAll() ([]karyawan.Response, error) {
	// Call Repo
	karyawanRepo, err := serv.Repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get karyawan data")
	}

	return karyawan.ToResponses(karyawanRepo), nil
}
