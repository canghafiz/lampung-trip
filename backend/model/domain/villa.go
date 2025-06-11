package domain

import "time"

type Villa struct {
	Id        int              `gorm:"primary_key;column:id;auto_increment"`
	Judul     string           `gorm:"column:judul"`
	Deskripsi string           `gorm:"column:deskripsi"`
	Rating    float64          `gorm:"column:rating"`
	Gambar    []GambarVilla    `gorm:"foreignKey:villa_id;references:id"`
	Fasilitas []FasilitasVilla `gorm:"foreignKey:villa_id;references:id"`
	Info      InfoVilla        `gorm:"foreignKey:villa_id;references:id"`
	Ulasan    []UlasanVilla    `gorm:"foreignKey:villa_id;references:id"`
}

func (Villa) TableName() string {
	return "villa"
}

type GambarVilla struct {
	Id      int    `gorm:"primary_key;column:id;auto_increment"`
	VillaId int    `gorm:"column:villa_id"`
	Url     string `gorm:"column:url"`
}

func (GambarVilla) TableName() string {
	return "gambarvilla"
}

type FasilitasVilla struct {
	Id      int    `gorm:"primary_key;column:id;auto_increment"`
	VillaId int    `gorm:"column:villa_id"`
	UrlIcon string `gorm:"column:url_icon"`
	Judul   string `gorm:"column:judul"`
}

func (FasilitasVilla) TableName() string {
	return "fasilitasvilla"
}

type InfoVilla struct {
	Id      int    `gorm:"primary_key;column:id;auto_increment"`
	VillaId int    `gorm:"column:villa_id"`
	Harga   string `gorm:"column:harga"`
	Lokasi  string `gorm:"column:lokasi"`
}

func (InfoVilla) TableName() string {
	return "infovilla"
}

type UlasanVilla struct {
	VillaId    int       `gorm:"primary_key;column:villa_id"`
	CustomerId int       `gorm:"primary_key;column:customer_id"`
	Rating     float64   `gorm:"column:rating"`
	Keterangan string    `gorm:"column:keterangan"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	Customer   Customer  `gorm:"foreignKey:customer_id;references:id"`
}

func (UlasanVilla) TableName() string {
	return "ulasanvilla"
}
