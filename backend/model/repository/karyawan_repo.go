package repository

import "lampung_trip/model/domain"

type KaryawanRepo interface {
	UpdateData(karyawan domain.Karyawan) error
	GetSingleByUserId(userId int) (*domain.Karyawan, error)
	GetAll() ([]domain.Karyawan, error)
}
