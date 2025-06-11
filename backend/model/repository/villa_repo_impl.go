package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"lampung_trip/model/domain"
)

type VillaRepoImpl struct {
	Db *gorm.DB
}

func NewVillaRepoImpl(db *gorm.DB) *VillaRepoImpl {
	return &VillaRepoImpl{Db: db}
}

func (repo *VillaRepoImpl) Create(villa domain.Villa) error {
	err := repo.Db.Create(&villa).Error
	if err != nil {
		return fmt.Errorf("failed to create villa: %w", err)
	}
	return nil
}

func (repo *VillaRepoImpl) Update(villa domain.Villa) error {
	tx := repo.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
		}
	}()

	// Handle Villa
	resultVilla := tx.Model(&domain.Villa{}).Where("id = ?", villa.Id).Updates(map[string]interface{}{
		"judul":     villa.Judul,
		"deskripsi": villa.Deskripsi,
	})
	if resultVilla.Error != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update villa: %w", resultVilla.Error)
	}

	// Handle GambarVilla
	if len(villa.Gambar) > 0 {
		if err := tx.Where("villa_id = ?", villa.Id).Delete(&domain.GambarVilla{}).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to delete existing gambar villa: %w", err)
		}
		for i := range villa.Gambar {
			villa.Gambar[i].VillaId = villa.Id // Pastikan villa_id terisi
		}
		if err := tx.Create(&villa.Gambar).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create new gambar villa: %w", err)
		}
	}

	// Handle FasilitasVilla
	if len(villa.Fasilitas) > 0 {
		if err := tx.Where("villa_id = ?", villa.Id).Delete(&domain.FasilitasVilla{}).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to delete existing fasilitas villa: %w", err)
		}
		for i := range villa.Fasilitas {
			villa.Fasilitas[i].VillaId = villa.Id // Pastikan villa_id terisi
		}
		if err := tx.Create(&villa.Fasilitas).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create new fasilitas villa: %w", err)
		}
	}

	// Handle InfoVilla
	if (domain.InfoVilla{}) != villa.Info {
		villa.Info.VillaId = villa.Id
		existingInfo := domain.InfoVilla{}
		if err := tx.Where("villa_id = ?", villa.Id).First(&existingInfo).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := tx.Create(&villa.Info).Error; err != nil {
					tx.Rollback()
					return fmt.Errorf("failed to create info villa: %w", err)
				}
			} else {
				tx.Rollback()
				return fmt.Errorf("failed to get existing info villa: %w", err)
			}
		} else {
			if err := tx.Model(&domain.InfoVilla{}).Where("villa_id = ?", villa.Id).Updates(villa.Info).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to update info villa: %w", err)
			}
		}
	}

	return nil
}

func (repo *VillaRepoImpl) GetAll() ([]domain.Villa, error) {
	var results []domain.Villa

	err := repo.Db.
		Select("id", "judul", "deskripsi", "rating").
		Preload("Gambar").
		Preload("Fasilitas").
		Preload("Info").
		Preload("Ulasan").
		Preload("Ulasan.Customer").
		Preload("Ulasan.Customer.User").
		Find(&results).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get all villas: %w", err)
	}

	return results, nil
}

func (repo *VillaRepoImpl) Delete(id int) error {
	tx := repo.Db.Begin()
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

	// Delete related data
	if err := tx.Where("villa_id = ?", id).Delete(&domain.GambarVilla{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete gambar villa: %w", err)
	}

	if err := tx.Where("villa_id = ?", id).Delete(&domain.FasilitasVilla{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete fasilitas villa: %w", err)
	}

	if err := tx.Where("villa_id = ?", id).Delete(&domain.InfoVilla{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete info villa: %w", err)
	}

	if err := tx.Where("villa_id = ?", id).Delete(&domain.UlasanVilla{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete ulasan villa: %w", err)
	}

	// Delete the main Villa
	result := tx.Delete(&domain.Villa{}, id)
	if result.Error != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete villa: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("no villa found with id: %d", id)
	}

	return nil
}
