package domain

import "time"

type Wisata struct {
	Id        int               `gorm:"primary_key;column:id;auto_increment"`
	Judul     string            `gorm:"judul"`
	Deskripsi string            `gorm:"deskripsi"`
	Rating    float64           `gorm:"rating"`
	Gambar    []GambarWisata    `gorm:"foreignKey:wisata_id;references:id"`
	Fasilitas []FasilitasWisata `gorm:"foreignKey:wisata_id;references:id"`
	Info      InfoWisata        `gorm:"foreignKey:wisata_id;references:id"`
	Ulasan    []UlasanWisata    `gorm:"foreignKey:wisata_id;references:id"`
}

func (Wisata) TableName() string {
	return "wisata"
}

type GambarWisata struct {
	Id       int    `gorm:"primary_key;column:id;auto_increment"`
	WisataId int    `gorm:"column:wisata_id"`
	Url      string `gorm:"column:url"`
}

func (GambarWisata) TableName() string {
	return "gambarwisata"
}

type FasilitasWisata struct {
	Id       int    `gorm:"primary_key;column:id;auto_increment"`
	WisataId int    `gorm:"column:wisata_id"`
	UrlIcon  string `gorm:"column:url_icon"`
	Judul    string `gorm:"column:judul"`
}

func (FasilitasWisata) TableName() string {
	return "fasilitaswisata"
}

type InfoWisata struct {
	Id         int    `gorm:"primary_key;column:id;auto_increment"`
	WisataId   int    `gorm:"column:wisata_id"`
	Harga      string `gorm:"column:harga"`
	WaktuBuka  string `gorm:"column:waktu_buka"`
	WaktuTutup string `gorm:"column:waktu_tutup"`
	HariBuka   string `gorm:"column:hari_buka"`
	HariTutup  string `gorm:"column:hari_tutup"`
	UrlLokasi  string `gorm:"column:url_lokasi"`
}

func (InfoWisata) TableName() string {
	return "infowisata"
}

type UlasanWisata struct {
	WisataId   int       `gorm:"column:wisata_id"`
	CustomerId int       `gorm:"column:customer_id"`
	Rating     float64   `gorm:"column:rating"`
	Keterangan string    `gorm:"column:keterangan"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	Customer   Customer  `gorm:"foreignKey:customer_id;references:id"`
}

func (UlasanWisata) TableName() string {
	return "ulasanwisata"
}
