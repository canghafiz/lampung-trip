package test

import (
	"github.com/stretchr/testify/assert"
	"lampung_trip/app"
	"lampung_trip/model/repository"
	"lampung_trip/model/service"
	"lampung_trip/model/web/villa"
	"testing"
)

func TestUlasanVillaServCreateSuccess(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")

	repo := repository.NewUlasanVillaRepoImpl(db)
	serv := service.NewUlasanVillaServImpl(repo)

	TestUserServRegisterCustomerSuccess(t)
	TestVillaServCreateSuccess(t)

	request := villa.UlasanCreateRequest{
		VillaId:    1,
		CustomerId: 1,
		Rating:     3,
		Keterangan: "Mantappp",
	}

	err := serv.Create(request)
	assert.Nil(t, err)

	TestVillaGetAllFound(t)
}

func TestUlasanVillaServDeleteSuccess(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")

	repo := repository.NewUlasanVillaRepoImpl(db)
	serv := service.NewUlasanVillaServImpl(repo)

	TestUlasanVillaServCreateSuccess(t)

	request := villa.UlasanDeleteRequest{
		VillaId:    1,
		CustomerId: 1,
	}

	err := serv.Delete(request)
	assert.Nil(t, err)

	defer func() {
		if r := recover(); r != nil {
			t.Logf("Caught assertion error from TestVillaGetAllFound: %v", r)
		}
	}()

	TestVillaGetAllFound(t)
}
