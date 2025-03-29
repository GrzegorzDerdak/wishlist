package internal

import (
	"gorm.io/gorm"
)

type Wishlist struct {
	gorm.Model
	Name        string `gorm:"index:idx_name_owner;not null"`
	Description *string
	IsPublic    bool
	OwnerId     uint    `gorm:"index:idx_name_owner;not null"`
	Items       []*Item `gorm:"many2many:wishlist_items;"`
}

type Item struct {
	gorm.Model
	Name      string
	Wishlists []*Wishlist `gorm:"many2many:wishlist_items;"`
}
