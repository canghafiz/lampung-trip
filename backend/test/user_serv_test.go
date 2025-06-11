package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"lampung_trip/app"
	"lampung_trip/model/repository"
	"lampung_trip/model/service"
	"lampung_trip/model/web/customer"
	"lampung_trip/model/web/karyawan"
	"lampung_trip/model/web/user"
	"testing"
)

func clearUserDb(db *gorm.DB) {
	db.Exec("DELETE FROM ulasanvilla")
	db.Exec("ALTER TABLE ulasanvilla AUTO_INCREMENT = 1")

	db.Exec("DELETE FROM ulasanwisata")
	db.Exec("ALTER TABLE ulasanwisata AUTO_INCREMENT = 1")

	db.Exec("DELETE FROM karyawan")
	db.Exec("ALTER TABLE karyawan AUTO_INCREMENT = 1")

	db.Exec("DELETE FROM customer")
	db.Exec("ALTER TABLE customer AUTO_INCREMENT = 1")

	db.Exec("DELETE FROM user")
	db.Exec("ALTER TABLE user AUTO_INCREMENT = 1")
}

func UserServLoginSuccess(t *testing.T, serv service.UserServ, username string, password string) {
	loginRequest := user.LoginRequest{
		Username: username,
		Password: password,
	}

	token, errLogin := serv.Login(loginRequest)
	assert.Nil(t, errLogin)
	assert.NotNil(t, token)

	fmt.Println(*token)
}

func TestUserServRegisterAdminSuccess(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")
	clearUserDb(db)

	repo := repository.NewUserRepoImpl(db)
	serv := service.NewUserServImpl(repo, "asdasdasd")

	request := user.RegisterAdminRequest{
		Username: "admin",
		Password: "123",
	}

	errAdmin := serv.RegisterAdmin(request)
	assert.Nil(t, errAdmin)

	UserServLoginSuccess(t, serv, "admin", "123")
}

func TestUserServUpdatePwSuccess(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")
	clearUserDb(db)

	repo := repository.NewUserRepoImpl(db)
	serv := service.NewUserServImpl(repo, "asdasdasd")

	TestUserServRegisterAdminSuccess(t)

	request := user.UpdatePwRequest{
		Username:    "admin",
		OldPassword: "123",
		NewPassword: "12345",
	}

	errUpdate := serv.UpdatePw(request)
	assert.Nil(t, errUpdate)

	UserServLoginSuccess(t, serv, "admin", "12345")
}

func TestUserServRegisterKaryawanSuccess(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")
	clearUserDb(db)

	repo := repository.NewUserRepoImpl(db)
	serv := service.NewUserServImpl(repo, "asdasdasd")

	request := user.RegisterKaryawanRequest{
		Username: "karyawan",
		Password: "123",
		Karyawan: karyawan.CreateRequest{
			Nama:  "Karyawan 1",
			Photo: "Ini foto karyawan",
			NoHp:  "0878017",
		},
	}

	errAdmin := serv.RegisterKaryawan(request)
	assert.Nil(t, errAdmin)

	UserServLoginSuccess(t, serv, "karyawan", "123")
}

func TestUserServRegisterCustomerSuccess(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")
	clearUserDb(db)

	repo := repository.NewUserRepoImpl(db)
	serv := service.NewUserServImpl(repo, "asdasdasd")

	request := user.RegisterCustomerRequest{
		Username: "cust",
		Password: "123",
		Customer: customer.CreateRequest{
			Nama:  "Cust 1",
			Photo: "Ini foto karyawan",
			NoHp:  "0821",
		},
	}

	errAdmin := serv.RegisterCustomer(request)
	assert.Nil(t, errAdmin)

	UserServLoginSuccess(t, serv, "cust", "123")
}

func TestUserServDeleteKaryawanSuccess(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")
	clearUserDb(db)

	repo := repository.NewUserRepoImpl(db)
	serv := service.NewUserServImpl(repo, "asdasdasd")

	TestUserServRegisterKaryawanSuccess(t)
	err := serv.DeleteKaryawan(1)
	assert.Nil(t, err)
}

func TestUserServDeleteAdminSuccess(t *testing.T) {
	db := app.OpenConnection("root", "", "localhost:3306", "db_lampung_trip")
	clearUserDb(db)

	repo := repository.NewUserRepoImpl(db)
	serv := service.NewUserServImpl(repo, "asdasdasd")

	TestUserServRegisterAdminSuccess(t)
	err := serv.DeleteAdmin(1)
	assert.Nil(t, err)
}
