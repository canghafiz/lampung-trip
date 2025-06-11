package repository

import (
	"fmt"
	"gorm.io/gorm"
	"lampung_trip/model/domain"
	"time"
)

type UlasanVillaRepoImpl struct {
	DB *gorm.DB
}

func NewUlasanVillaRepoImpl(DB *gorm.DB) *UlasanVillaRepoImpl {
	return &UlasanVillaRepoImpl{DB: DB}
}

func (repo *UlasanVillaRepoImpl) Create(ulasan domain.UlasanVilla) error {
	ulasan.CreatedAt = time.Now()
	err := repo.DB.Create(&ulasan).Error
	if err != nil {
		return fmt.Errorf("failed to create ulasan villa: %w", err)
	}

	return nil
}

func (repo *UlasanVillaRepoImpl) Delete(ulasan domain.UlasanVilla) error {
	result := repo.DB.Where("villa_id = ? AND customer_id = ?", ulasan.VillaId, ulasan.CustomerId).Delete(&domain.UlasanVilla{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete ulasan villa: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("ulasan villa tidak ditemukan dengan VillaId = %d dan CustomerId = %d", ulasan.VillaId, ulasan.CustomerId)
	}
	return nil
}
