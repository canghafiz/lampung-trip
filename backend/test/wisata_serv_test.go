package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"lampung_trip/app"
	"lampung_trip/model/repository"
	"lampung_trip/model/service"
	"lampung_trip/model/web/wisata"
	"testing"
)

func clearWisataDb(db *gorm.DB) {
	db.Exec("DELETE FROM ulasanwisata")
	db.Exec("ALTER TABLE ulasanwisata AUTO_INCREMENT = 1")

	db.Exec("DELETE FROM gambarwisata")
	db.Exec("ALTER TABLE gambarwisata AUTO_INCREMENT = 1")

	db.Exec("DELETE FROM fasilitaswisata")
	db.Exec("ALTER TABLE fasilitaswisata AUTO_INCREMENT = 1")

	db.Exec("DELETE FROM infowisata")
	db.Exec("ALTER TABLE infowisata AUTO_INCREMENT = 1")

	db.Exec("DELETE FROM wisata")
	db.Exec("ALTER TABLE wisata AUTO_INCREMENT = 1")
}

func TestWisataServCreateSuccess(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")
	clearWisataDb(db)

	repo := repository.NewWisataRepoImpl(db)
	serv := service.NewWisataServImpl(repo)

	request := wisata.CreateRequest{
		Judul:     "Wisata",
		Deskripsi: "Dadafdasdasdasfasdf",
		Info: wisata.InfoRequest{
			Harga:      "20000",
			HariBuka:   "Senin",
			HariTutup:  "Minggu",
			UrlLokasi:  "Lokasi",
			WaktuBuka:  "07:00",
			WaktuTutup: "17:00",
		},
		Fasilitas: []wisata.FasilitasRequest{
			wisata.FasilitasRequest{
				Judul:   "Wisata",
				UrlIcon: "dasdsd",
			},
		},
		Gambar: []wisata.GambarRequest{
			wisata.GambarRequest{
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

func TestWisataGetAllFound(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")
	repo := repository.NewWisataRepoImpl(db)
	serv := service.NewWisataServImpl(repo)

	results, _ := serv.GetAll()
	assert.Equal(t, 1, len(results))

	fmt.Println(results)
}

func TestWisataServUpdateSuccess(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")

	TestWisataServCreateSuccess(t)

	repo := repository.NewWisataRepoImpl(db)
	serv := service.NewWisataServImpl(repo)

	request := wisata.UpdateRequest{
		Id:        1,
		Judul:     "Wisata Update",
		Deskripsi: "Dadafdasdasdasfasdf",
		Info: wisata.InfoRequest{
			Harga:      "20000",
			HariBuka:   "Senin",
			HariTutup:  "Minggu",
			UrlLokasi:  "Lokasi",
			WaktuBuka:  "07:00",
			WaktuTutup: "17:00",
		},
		Fasilitas: []wisata.FasilitasRequest{
			wisata.FasilitasRequest{
				Judul:   "Wisata asdasdsa",
				UrlIcon: "aa",
			},
		},
		Gambar: []wisata.GambarRequest{
			wisata.GambarRequest{
				Url: "Lokasiii",
			},
		},
	}

	err := serv.Update(request)
	assert.Nil(t, err)
}

func TestWisataServDeleteSuccess(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")

	TestWisataServCreateSuccess(t)

	repo := repository.NewWisataRepoImpl(db)
	serv := service.NewWisataServImpl(repo)

	err := serv.Delete(1)
	assert.Nil(t, err)
}
