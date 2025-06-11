package service

import (
	"fmt"
	"lampung_trip/helper"
	"lampung_trip/model/repository"
	"lampung_trip/model/web/user"
)

type UserServImpl struct {
	Repo      repository.UserRepo
	SecretKey string
}

func NewUserServImpl(repo repository.UserRepo, secretKey string) *UserServImpl {
	return &UserServImpl{Repo: repo, SecretKey: secretKey}
}

func (serv *UserServImpl) LoginAdminKaryawan(request user.LoginRequest) (*string, error) {
	// Call Repo
	repoResult, _ := serv.Repo.GetSingleByUsername(request.Username)

	// When user not found
	if repoResult == nil {
		return nil, fmt.Errorf("user not found")
	}

	// User role must be admin or karyawan
	if !(repoResult.Role == "Admin" || repoResult.Role == "Karyawan") {
		return nil, fmt.Errorf("not allowed")
	}

	// Check Password
	passwordValid := helper.CheckPassword(repoResult.Password, request.Password)
	if !passwordValid {
		return nil, fmt.Errorf("password wrong")
	}

	// Generate jwt token
	token, errToken := helper.GenerateJWT(repoResult.Id, repoResult.Role, serv.SecretKey)
	if errToken != nil {
		return nil, fmt.Errorf("failed to generate token")
	}

	return &token, nil
}

func (serv *UserServImpl) Login(request user.LoginRequest) (*string, error) {
	// Call Repo
	repoResult, _ := serv.Repo.GetSingleByUsername(request.Username)

	// when user not found
	if repoResult == nil {
		return nil, fmt.Errorf("user not found")
	}

	// User role must be admin or karyawan
	if !(repoResult.Role == "Customer") {
		return nil, fmt.Errorf("user not found")
	}

	// Check Password
	passwordValid := helper.CheckPassword(repoResult.Password, request.Password)
	if !passwordValid {
		return nil, fmt.Errorf("password wrong")
	}

	// Generate jwt token
	token, errToken := helper.GenerateJWT(repoResult.Id, repoResult.Role, serv.SecretKey)
	if errToken != nil {
		return nil, fmt.Errorf("failed to generate token")
	}

	return &token, nil
}

func (serv *UserServImpl) UpdatePw(request user.UpdatePwRequest) error {
	// Call Repo
	repoResult, _ := serv.Repo.GetSingleByUsername(request.Username)

	// when user not found
	if repoResult == nil {
		return fmt.Errorf("user not found")
	}

	// Check Password
	passwordValid := helper.CheckPassword(repoResult.Password, request.OldPassword)
	if !passwordValid {
		return fmt.Errorf("password wrong")
	}

	// Update
	domain := user.UpdatePwRequestToDomain(request, repoResult.Id)
	errUpdate := serv.Repo.UpdatePw(domain)
	if errUpdate != nil {
		return fmt.Errorf("failed to update password")
	}

	return nil
}

func (serv *UserServImpl) RegisterAdmin(request user.RegisterAdminRequest) error {
	// Call Repo
	repoResult, _ := serv.Repo.GetSingleByUsername(request.Username)

	// when user found
	if repoResult != nil {
		return fmt.Errorf("user already exists")
	}

	// Create
	domain := user.RegisterAdminRequestToDomain(request)
	errCreate := serv.Repo.Create(domain)
	if errCreate != nil {
		return fmt.Errorf("failed to create admin")
	}

	return nil
}

func (serv *UserServImpl) RegisterKaryawan(request user.RegisterKaryawanRequest) error {
	// Call Repo
	repoResult, _ := serv.Repo.GetSingleByUsername(request.Username)

	// when user found
	if repoResult != nil {
		return fmt.Errorf("user already exists")
	}

	// Create
	domain := user.RegisterKaryawanRequestToDomain(request)
	errCreate := serv.Repo.Create(domain)
	if errCreate != nil {
		return fmt.Errorf("failed to create user")
	}

	return nil
}

func (serv *UserServImpl) RegisterCustomer(request user.RegisterCustomerRequest) error {
	// Call Repo
	repoResult, _ := serv.Repo.GetSingleByUsername(request.Username)

	// when user found
	if repoResult != nil {
		return fmt.Errorf("user already exists")
	}

	// Create
	domain := user.RegisterCustomerRequestToDomain(request)
	errCreate := serv.Repo.Create(domain)
	if errCreate != nil {
		return fmt.Errorf("failed to create user")
	}

	return nil
}

func (serv *UserServImpl) DeleteKaryawan(id int) error {
	// Call Repo
	err := serv.Repo.DeleteKaryawan(id)
	if err != nil {
		return fmt.Errorf("failed to delete karyawan")
	}

	return nil
}

func (serv *UserServImpl) DeleteAdmin(id int) error {
	err := serv.Repo.DeleteAdmin(id)
	if err != nil {
		return fmt.Errorf("failed to delete admin")
	}

	return nil
}
