package repository

import (
	"fmt"
	"gorm.io/gorm"
	"lampung_trip/model/domain"
)

type UserRepoImpl struct {
	DB *gorm.DB
}

func NewUserRepoImpl(DB *gorm.DB) *UserRepoImpl {
	return &UserRepoImpl{DB: DB}
}

func (repo *UserRepoImpl) Create(user domain.User) error {
	err := repo.DB.Create(&user).Error
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (repo *UserRepoImpl) GetSingleByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := repo.DB.
		Select("id", "username", "password", "role").
		Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}
	return &user, nil
}

func (repo *UserRepoImpl) UpdatePw(user domain.User) error {
	err := repo.DB.
		Model(&user).
		Where("id = ?", user.Id).
		Updates(map[string]interface{}{
			"password": user.Password,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update password")
	}

	return nil
}

func (repo *UserRepoImpl) DeleteKaryawan(id int) error {
	// Delete relate id from karyawan table
	result := repo.DB.Where("user_id = ?", id).Delete(&domain.Karyawan{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete related karyawan data: %w", result.Error)
	}

	// Then delete user from user table
	err := repo.DB.Delete(&domain.User{}, id).Error
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

func (repo *UserRepoImpl) DeleteAdmin(id int) error {
	err := repo.DB.Delete(&domain.User{
		Role: "Admin",
	}, id).Error
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
