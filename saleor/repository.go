package saleor

import "gorm.io/gorm"

type SaleorConfigRepository struct {
	DB *gorm.DB
}

func NewSaleorConfigRepository(db *gorm.DB) *SaleorConfigRepository {
	return &SaleorConfigRepository{DB: db}
}

func (r *SaleorConfigRepository) RegisterSaleorDomain(config *SaleorConfig) (*SaleorConfig, error) {
	var existing SaleorConfig

	if err := r.DB.Where("domain = ?", config.Domain).First(&existing).Error; err == nil {
		existing.ApiUrl = config.ApiUrl
		existing.SchemaVersion = config.SchemaVersion
		existing.AuthToken = config.AuthToken
		if err := r.DB.Save(&existing).Error; err != nil {
			return nil, err
		}
		return &existing, nil
	} else if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if err := r.DB.Create(config).Error; err != nil {
		return nil, err
	}

	return config, nil
}

func (r *SaleorConfigRepository) GetConfigByDomain(domain string) (*SaleorConfig, error) {
	var config SaleorConfig
	if err := r.DB.Where("domain = ?", domain).First(&config).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // No record found
		}
		return nil, err
	}
	return &config, nil
}
