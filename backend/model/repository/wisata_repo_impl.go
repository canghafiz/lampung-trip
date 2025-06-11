package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"lampung_trip/model/domain"
)

type WisataRepoImpl struct {
	DB *gorm.DB
}

func NewWisataRepoImpl(DB *gorm.DB) *WisataRepoImpl {
	return &WisataRepoImpl{DB: DB}
}

func (repo *WisataRepoImpl) Create(wisata domain.Wisata) error {
	err := repo.DB.Create(&wisata).Error
	if err != nil {
		return fmt.Errorf("failed to create wisata: %w", err)
	}
	return nil
}

func (repo *WisataRepoImpl) Update(wisata domain.Wisata) error {
	tx := repo.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
		}
	}()

	// Handle Wisata
	resultWisata := tx.Model(&domain.Wisata{}).Where("id = ?", wisata.Id).Updates(map[string]interface{}{
		"judul":     wisata.Judul,
		"deskripsi": wisata.Deskripsi,
	})
	if resultWisata.Error != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update wisata: %w", resultWisata.Error)
	}
	// GambarWisata handle
	if len(wisata.Gambar) > 0 {
		if err := tx.Where("wisata_id = ?", wisata.Id).Delete(&domain.GambarWisata{}).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to delete existing gambar wisata: %w", err)
		}
		for i := range wisata.Gambar {
			wisata.Gambar[i].WisataId = wisata.Id // Pastikan wisata_id terisi
		}
		if err := tx.Create(&wisata.Gambar).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create new gambar wisata: %w", err)
		}
	}

	// Handle FasilitasWisata
	if len(wisata.Fasilitas) > 0 {
		if err := tx.Where("wisata_id = ?", wisata.Id).Delete(&domain.FasilitasWisata{}).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to delete existing fasilitas wisata: %w", err)
		}
		for i := range wisata.Fasilitas {
			wisata.Fasilitas[i].WisataId = wisata.Id // Pastikan wisata_id terisi
		}
		if err := tx.Create(&wisata.Fasilitas).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create new fasilitas wisata: %w", err)
		}
	}

	// Handle InfoWisata
	if (domain.InfoWisata{}) != wisata.Info {
		wisata.Info.WisataId = wisata.Id
		existingInfo := domain.InfoWisata{}
		if err := tx.Where("wisata_id = ?", wisata.Id).First(&existingInfo).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := tx.Create(&wisata.Info).Error; err != nil {
					tx.Rollback()
					return fmt.Errorf("failed to create info wisata: %w", err)
				}
			} else {
				tx.Rollback()
				return fmt.Errorf("failed to get existing info wisata: %w", err)
			}
		} else {
			if err := tx.Model(&domain.InfoWisata{}).Where("wisata_id = ?", wisata.Id).Updates(wisata.Info).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to update info wisata: %w", err)
			}
		}
	}

	return nil
}

func (repo *WisataRepoImpl) GetAll() ([]domain.Wisata, error) {
	var results []domain.Wisata

	err := repo.DB.
		Select("id", "judul", "deskripsi", "rating").
		Preload("Gambar").
		Preload("Fasilitas").
		Preload("Info").
		Preload("Ulasan").
		Preload("Ulasan.Customer").
		Preload("Ulasan.Customer.User").
		Find(&results).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get all wisata: %w", err)
	}

	return results, nil
}

func (repo *WisataRepoImpl) Delete(id int) error {
	tx := repo.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			// Log the commit error
		}
	}()

	// Delete constraint
	if err := tx.Where("wisata_id = ?", id).Delete(&domain.GambarWisata{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete gambar wisata: %w", err)
	}

	if err := tx.Where("wisata_id = ?", id).Delete(&domain.FasilitasWisata{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete fasilitas wisata: %w", err)
	}

	if err := tx.Where("wisata_id = ?", id).Delete(&domain.InfoWisata{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete info wisata: %w", err)
	}

	if err := tx.Where("wisata_id = ?", id).Delete(&domain.UlasanWisata{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete ulasan wisata: %w", err)
	}

	// Delete Main Wisata
	result := tx.Delete(&domain.Wisata{}, id)
	if result.Error != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete wisata: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("no wisata found with id: %d", id)
	}

	return nil
}
