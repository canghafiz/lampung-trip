package domain

type User struct {
	Id       int      `gorm:"primary_key;column:id;auto_increment"`
	Username string   `gorm:"column:username"`
	Password string   `gorm:"column:password"`
	Role     string   `gorm:"column:role"`
	Customer Customer `gorm:"foreignKey:user_id;references:id"`
	Karyawan Karyawan `gorm:"foreignKey:user_id;references:id"`
}

func (User) TableName() string {
	return "user"
}
