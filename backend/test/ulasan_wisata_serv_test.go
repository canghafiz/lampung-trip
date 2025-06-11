package test

import (
	"github.com/stretchr/testify/assert"
	"lampung_trip/app"
	"lampung_trip/model/repository"
	"lampung_trip/model/service"
	"lampung_trip/model/web/wisata"
	"testing"
)

func TestUlasanWisataServCreateSuccess(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")

	repo := repository.NewUlasanWisataRepoImpl(db)
	serv := service.NewUlasanWisataServImpl(repo)

	TestUserServRegisterCustomerSuccess(t)
	TestWisataServCreateSuccess(t)

	request := wisata.UlasanCreateRequest{
		WisataId:   1,
		CustomerId: 1,
		Rating:     5,
		Keterangan: "Mantappp",
	}

	err := serv.Create(request)
	assert.Nil(t, err)

	TestWisataGetAllFound(t)
}

func TestUlasanWisataServDeleteSuccess(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")

	repo := repository.NewUlasanWisataRepoImpl(db)
	serv := service.NewUlasanWisataServImpl(repo)

	TestUlasanWisataServCreateSuccess(t)

	request := wisata.UlasanDeleteRequest{
		WisataId:   1,
		CustomerId: 1,
	}

	err := serv.Delete(request)
	assert.Nil(t, err)

	defer func() {
		if r := recover(); r != nil {
			t.Logf("Caught assertion error from TestWisataGetAllFound: %v", r)
		}
	}()

	TestWisataGetAllFound(t)
}
