package domain

type Karyawan struct {
	Id     int    `gorm:"primary_key;column:id;auto_increment"`
	UserId int    `gorm:"column:user_id"`
	Nama   string `gorm:"column:nama"`
	Photo  string `gorm:"column:photo"`
	NoHp   string `gorm:"column:no_hp"`
	User   *User  `gorm:"foreignKey:user_id;references:id"`
}

func (Karyawan) TableName() string {
	return "karyawan"
}
