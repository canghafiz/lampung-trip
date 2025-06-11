package service

import "lampung_trip/model/web/user"

type UserServ interface {
	LoginAdminKaryawan(request user.LoginRequest) (*string, error)
	Login(request user.LoginRequest) (*string, error)
	UpdatePw(request user.UpdatePwRequest) error
	RegisterAdmin(request user.RegisterAdminRequest) error
	RegisterKaryawan(request user.RegisterKaryawanRequest) error
	RegisterCustomer(request user.RegisterCustomerRequest) error
	DeleteKaryawan(id int) error
	DeleteAdmin(id int) error
}
