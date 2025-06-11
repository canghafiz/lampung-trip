package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"lampung_trip/app"
	"lampung_trip/model/repository"
	"lampung_trip/model/service"
	"lampung_trip/model/web/karyawan"
	"testing"
)

func TestKaryawanServUpdateData(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")
	clearUserDb(db)

	repo := repository.NewKaryawanRepoImpl(db)
	serv := service.NewKaryawanServImpl(repo)

	TestUserServRegisterKaryawanSuccess(t)

	request := karyawan.UpdateDataRequest{
		Id:    1,
		Nama:  "hafiz",
		Photo: "",
		NoHp:  "0821",
	}

	errUpdate := serv.UpdateData(request)
	assert.Nil(t, errUpdate)
}

func TestKaryawanServGetSingleByUserId(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")
	clearUserDb(db)

	repo := repository.NewKaryawanRepoImpl(db)
	serv := service.NewKaryawanServImpl(repo)

	TestUserServRegisterKaryawanSuccess(t)

	result, errResult := serv.GetSingleByUserId(1)
	assert.Nil(t, errResult)
	assert.NotNil(t, result)

	fmt.Println(result)
}

func TestKaryawanServGetAllFound(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")
	clearUserDb(db)

	repo := repository.NewKaryawanRepoImpl(db)
	serv := service.NewKaryawanServImpl(repo)

	TestUserServRegisterKaryawanSuccess(t)

	result, errResult := serv.GetAll()
	assert.Nil(t, errResult)
	assert.NotNil(t, result)

	fmt.Println(result)
}
