package internal

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Handle layer
type WishlistHandler struct {
	wishlistService *WishlistService
}

func NewWishlistHandler(wishlistService *WishlistService) *WishlistHandler {
	return &WishlistHandler{
		wishlistService: wishlistService,
	}
}

type CreateWishlistRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=255"`
	Description string `json:"description" validate:"max=500"`
	IsPublic    bool   `json:"isPublic"`
	Items       []Item `json:"items"`
}

func (h *WishlistHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var req CreateWishlistRequest

	// Decode the request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the request
	if err := validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wishlist := Wishlist{
		Name:        req.Name,
		Description: &req.Description,
		IsPublic:    req.IsPublic,
	}
	// Create the wishlist
	if err := h.wishlistService.Create(&wishlist); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the wishlist to the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wishlist)
}

func (h *WishlistHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// Parse the request URL
	id := r.URL.Query().Get("id")

	// Get the wishlist by ID
	wishlist, err := h.wishlistService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the wishlist to the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wishlist)
}
