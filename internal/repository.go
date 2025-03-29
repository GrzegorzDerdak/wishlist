package internal

import (
	"gorm.io/gorm"
)

type WishlistRepository struct {
	DB *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) *WishlistRepository {
	return &WishlistRepository{
		DB: db,
	}
}

func (r *WishlistRepository) Create(wishlist *Wishlist) error {
	return r.DB.Create(wishlist).Error
}

func (r *WishlistRepository) GetByID(id string) (*Wishlist, error) {
	var wishlist Wishlist
	if err := r.DB.First(&wishlist, id).Error; err != nil {
		return nil, err
	}
	return &wishlist, nil
}
