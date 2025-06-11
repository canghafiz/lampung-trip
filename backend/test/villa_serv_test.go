package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"lampung_trip/app"
	"lampung_trip/model/repository"
	"lampung_trip/model/service"
	"lampung_trip/model/web/villa"
	"testing"
)

func clearVillaDb(db *gorm.DB) {
	db.Exec("DELETE FROM ulasanvilla")
	db.Exec("ALTER TABLE ulasanvilla AUTO_INCREMENT = 1")

	db.Exec("DELETE FROM gambarvilla")
	db.Exec("ALTER TABLE gambarvilla AUTO_INCREMENT = 1")

	db.Exec("DELETE FROM fasilitasvilla")
	db.Exec("ALTER TABLE fasilitasvilla AUTO_INCREMENT = 1")

	db.Exec("DELETE FROM infovilla")
	db.Exec("ALTER TABLE infovilla AUTO_INCREMENT = 1")

	db.Exec("DELETE FROM villa")
	db.Exec("ALTER TABLE villa AUTO_INCREMENT = 1")
}

func TestVillaServCreateSuccess(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")
	clearVillaDb(db)

	repo := repository.NewVillaRepoImpl(db)
	serv := service.NewVillaServImpl(repo)

	request := villa.CreateRequest{
		Judul:     "Villa",
		Deskripsi: "Dadafdasdasdasfasdf",
		Info: villa.InfoRequest{
			Harga:  "20000",
			Lokasi: "Lokasi",
		},
		Fasilitas: []villa.FasilitasRequest{
			villa.FasilitasRequest{
				Judul:   "Villa",
				UrlIcon: "dasdsd",
			},
		},
		Gambar: []villa.GambarRequest{
			villa.GambarRequest{
				Url: "Lokasi",
			},
		},
	}

	err := serv.Create(request)
	assert.Nil(t, err)

	results, _ := serv.GetAll()
	assert.Equal(t, 1, len(results))

	fmt.Println(results)
}

func TestVillaGetAllFound(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")
	repo := repository.NewVillaRepoImpl(db)
	serv := service.NewVillaServImpl(repo)

	results, _ := serv.GetAll()
	assert.Equal(t, 1, len(results))

	fmt.Println(results)
}

func TestVillaServUpdateSuccess(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")

	TestVillaServCreateSuccess(t)

	repo := repository.NewVillaRepoImpl(db)
	serv := service.NewVillaServImpl(repo)

	request := villa.UpdateRequest{
		Id:        1,
		Judul:     "Villa Update",
		Deskripsi: "Dadafdasdasdasfasdf",
		Info: villa.InfoRequest{
			Harga:  "20000",
			Lokasi: "Update",
		},
		Fasilitas: []villa.FasilitasRequest{
			villa.FasilitasRequest{
				Judul:   "Villa asdasdsa",
				UrlIcon: "aa",
			},
		},
		Gambar: []villa.GambarRequest{
			villa.GambarRequest{
				Url: "Lokasiii",
			},
		},
	}

	err := serv.Update(request)
	assert.Nil(t, err)
}

func TestVillaServDeleteSuccess(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")

	TestVillaServCreateSuccess(t)

	repo := repository.NewVillaRepoImpl(db)
	serv := service.NewVillaServImpl(repo)

	err := serv.Delete(1)
	assert.Nil(t, err)
}
