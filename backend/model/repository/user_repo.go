package repository

import "lampung_trip/model/domain"

type UserRepo interface {
	Create(user domain.User) error
	GetSingleByUsername(username string) (*domain.User, error)
	UpdatePw(user domain.User) error
	DeleteKaryawan(id int) error
	DeleteAdmin(id int) error
}
