package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"lampung_trip/model/domain"
)

type CustomerRepoImpl struct {
	DB *gorm.DB
}

func NewCustomerRepoImpl(DB *gorm.DB) *CustomerRepoImpl {
	return &CustomerRepoImpl{DB: DB}
}

func (repo *CustomerRepoImpl) UpdateData(customer domain.Customer) error {
	err := repo.DB.
		Model(&customer).
		Where("user_id = ?", customer.UserId).
		Updates(map[string]interface{}{
			"nama":  customer.Nama,
			"photo": customer.Photo,
			"no_hp": customer.NoHp,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update data customer")
	}

	return nil
}

func (repo *CustomerRepoImpl) GetSingleByUserId(userId int) (*domain.Customer, error) {
	var result domain.Customer

	err := repo.DB.
		Select("id", "user_id", "nama", "photo", "no_hp").
		Preload("User").
		Where("user_id = ?", userId).
		First(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("customer not found")
		}
		return nil, fmt.Errorf("failed to get customer by user ID: %w", err)
	}

	return &result, nil
}

func (repo *CustomerRepoImpl) GetAll() ([]domain.Customer, error) {
	var results []domain.Customer

	err := repo.DB.
		Select("id", "user_id", "nama", "photo", "no_hp").
		Preload("User").
		Find(&results).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("customer not found")
		}
		return nil, fmt.Errorf("failed to get customer")
	}

	return results, nil
}
