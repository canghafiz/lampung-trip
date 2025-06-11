package repository

import (
	"fmt"
	"gorm.io/gorm"
	"lampung_trip/model/domain"
	"time"
)

type UlasanWisataRepoImpl struct {
	DB *gorm.DB
}

func NewUlasanWisataRepoImpl(DB *gorm.DB) *UlasanWisataRepoImpl {
	return &UlasanWisataRepoImpl{DB: DB}
}

func (repo *UlasanWisataRepoImpl) Create(ulasan domain.UlasanWisata) error {
	ulasan.CreatedAt = time.Now()
	err := repo.DB.Create(&ulasan).Error
	if err != nil {
		return fmt.Errorf("failed to create ulasan wisata: %w", err)
	}

	return nil
}

func (repo *UlasanWisataRepoImpl) Delete(ulasan domain.UlasanWisata) error {
	result := repo.DB.Where("wisata_id = ? AND customer_id = ?", ulasan.WisataId, ulasan.CustomerId).Delete(&domain.UlasanWisata{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete ulasan wisata: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("ulasan wisata tidak ditemukan dengan WisataId = %d dan CustomerId = %d", ulasan.WisataId, ulasan.CustomerId)
	}
	return nil
}
