package wishlists

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Wishlist struct {
	gorm.Model
	Name        string `gorm:"index:idx_name_owner;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Description string
	IsPublic    bool
	OwnerId     string  `gorm:"index:idx_name_owner;not null"`
	Items       []*Item `gorm:"many2many:wishlist_items;"`
}

type Item struct {
	gorm.Model
	Name       string
	Wishlists  []*Wishlist `gorm:"many2many:wishlist_items;"`
	BusinessID string
}

type CreateWishlistRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=255"`
	Description string `json:"description" validate:"max=500"`
	IsPublic    bool   `json:"isPublic"`
	Items       []Item `json:"items"`
}
type UpdateWishlistRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=255"`
	Description string `json:"description" validate:"max=500"`
	IsPublic    bool   `json:"isPublic"`
}

type WishlistResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
	Description string `json:"description"`
	IsPublic    bool   `json:"isPublic"`
	Items       []Item `json:"items"`
}

func CreateSingleWishlistResponse(wishlist *Wishlist) WishlistResponse {
	var items []Item
	for _, itemPtr := range wishlist.Items {
		if itemPtr != nil {
			items = append(items, *itemPtr)
		}
	}

	return WishlistResponse{
		ID:          fmt.Sprintf("%d", wishlist.ID),
		Name:        wishlist.Name,
		CreatedAt:   wishlist.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   wishlist.UpdatedAt.Format(time.RFC3339),
		Description: wishlist.Description,
		IsPublic:    wishlist.IsPublic,
		Items:       items,
	}
}

func CreateAllWishlistsResponse(wishlists []*Wishlist) []WishlistResponse {
	var responses []WishlistResponse
	for _, wishlist := range wishlists {
		responses = append(responses, CreateSingleWishlistResponse(wishlist))
	}
	return responses
}
