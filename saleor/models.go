package saleor

import (
	"errors"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type SaleorConfig struct {
	gorm.Model
	Domain        string `gorm:"size:255;not null"`
	ApiUrl        string `gorm:"size:255;not null"`
	SchemaVersion string `gorm:"size:50;not null"`
	AuthToken     string `gorm:"size:512"` // Increased size for tokens
}

func (c SaleorConfig) Validate() error {
	if c.Domain == "" {
		return errors.New("Saleor-Domain header is required")
	}

	if c.ApiUrl == "" {
		return errors.New("Saleor-Api-Url header is required")
	}

	if !strings.HasPrefix(c.ApiUrl, "http") {
		return errors.New("Saleor-Api-Url must be a valid URL")
	}

	return nil
}

func (c SaleorConfig) ToConfig() SaleorConfig {
	return SaleorConfig{
		Domain:        c.Domain,
		ApiUrl:        c.ApiUrl,
		SchemaVersion: c.SchemaVersion,
		AuthToken:     c.AuthToken,
	}
}

func ParseSaleorHeaders(r *http.Request) SaleorConfig {
	return SaleorConfig{
		Domain:        r.Header.Get("Saleor-Domain"),
		ApiUrl:        r.Header.Get("Saleor-Api-Url"),
		SchemaVersion: r.Header.Get("Saleor-Schema-Version"),
	}
}
