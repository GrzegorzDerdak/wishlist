package wishlists

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

func (r *WishlistRepository) GetAllByUserID(userID string) ([]*Wishlist, error) {
	var wishlists []*Wishlist
	if err := r.DB.Where("owner_id = ?", userID).Find(&wishlists).Error; err != nil {
		return nil, err
	}
	return wishlists, nil
}

func (r *WishlistRepository) Update(wishlist *Wishlist) (*Wishlist, error) {
	var existingWishlist Wishlist
	if err := r.DB.First(&existingWishlist, wishlist.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	if existingWishlist.OwnerId != wishlist.OwnerId {
		return nil, gorm.ErrRecordNotFound // or a custom error indicating permission denied
	}

	existingWishlist.Name = wishlist.Name
	existingWishlist.Description = wishlist.Description
	existingWishlist.IsPublic = wishlist.IsPublic
	existingWishlist.Items = wishlist.Items

	if err := r.DB.Save(&existingWishlist).Error; err != nil {
		return nil, err
	}

	return &existingWishlist, nil
}

func (r *WishlistRepository) Delete(id string) error {
	var wishlist Wishlist
	if err := r.DB.First(&wishlist, id).Error; err != nil {
		return err
	}
	return r.DB.Delete(&wishlist).Error
}
