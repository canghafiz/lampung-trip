package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"lampung_trip/model/domain"
)

type KaryawanRepoImpl struct {
	DB *gorm.DB
}

func NewKaryawanRepoImpl(DB *gorm.DB) *KaryawanRepoImpl {
	return &KaryawanRepoImpl{DB: DB}
}

func (repo *KaryawanRepoImpl) UpdateData(karyawan domain.Karyawan) error {
	err := repo.DB.
		Model(&karyawan).
		Where("user_id = ?", karyawan.UserId).
		Updates(map[string]interface{}{
			"nama":  karyawan.Nama,
			"photo": karyawan.Photo,
			"no_hp": karyawan.NoHp,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update data karyawan")
	}

	return nil
}

func (repo *KaryawanRepoImpl) GetSingleByUserId(userId int) (*domain.Karyawan, error) {
	var result domain.Karyawan

	err := repo.DB.
		Select("id", "user_id", "nama", "photo", "no_hp").
		Preload("User").
		Where("user_id = ?", userId).
		First(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("karyawan not found")
		}
		return nil, fmt.Errorf("failed to get karyawan by user ID: %w", err)
	}

	return &result, nil
}

func (repo *KaryawanRepoImpl) GetAll() ([]domain.Karyawan, error) {
	var results []domain.Karyawan

	err := repo.DB.
		Select("id", "user_id", "nama", "photo", "no_hp").
		Preload("User").
		Find(&results).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("karyawan not found")
		}
		return nil, fmt.Errorf("failed to get karyawan")
	}

	return results, nil
}
