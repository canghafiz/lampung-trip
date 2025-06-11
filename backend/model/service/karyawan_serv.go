package service

import "lampung_trip/model/web/karyawan"

type KaryawanServ interface {
	UpdateData(request karyawan.UpdateDataRequest) error
	GetSingleByUserId(userId int) (*karyawan.Response, error)
	GetAll() ([]karyawan.Response, error)
}
