package wishlists

import (
	"encoding/json"

	"net/http"
	"strconv"

	"wishlist/internal"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

var validate = validator.New()

type WishlistHandler struct {
	wishlistService *WishlistService
}

func NewWishlistHandler(wishlistService *WishlistService) *WishlistHandler {
	return &WishlistHandler{
		wishlistService: wishlistService,
	}
}

func (h *WishlistHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(internal.UserIDCtxKey).(string)
	l := zerolog.Ctx(r.Context())

	var req CreateWishlistRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		l.Warn().Msg("Failed to decode request body")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {
		l.Warn().Msg("Invalid request payload")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wishlist := Wishlist{
		Name:        req.Name,
		Description: req.Description,
		IsPublic:    req.IsPublic,
		OwnerId:     userID,
	}

	if err := h.wishlistService.Create(&wishlist); err != nil {
		l.Err(err).Msg("Failed to create wishlist")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	l.Debug().Int("WishlistID", int(wishlist.ID)).Msg("Created wishlist")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateSingleWishlistResponse(&wishlist))
}

func (h *WishlistHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(internal.UserIDCtxKey).(string)
	l := zerolog.Ctx(r.Context())

	wishlists, err := h.wishlistService.GetAllByUserID(userID)
	if err != nil {
		l.Err(err).Msg("Failed to retrieve wishlists")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	l.Debug().Int("Count", len(wishlists)).Msg("Retrieved wishlists")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CreateAllWishlistsResponse(wishlists))
}

func (h *WishlistHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(internal.UserIDCtxKey).(string)
	params := mux.Vars(r)
	wishlistId := params["wishlist_id"]
	l := zerolog.Ctx(r.Context())

	wishlist, err := h.wishlistService.GetByID(wishlistId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if wishlist == nil {
		l.Warn().Msg("Wishlist not found")
		http.Error(w, "Wishlist not found", http.StatusNotFound)
		return
	}

	if wishlist.OwnerId != userID {
		l.Warn().Msg("User attempted to access a wishlist they do not own")
		http.Error(w, "Forbidden: You do not have access to this wishlist", http.StatusForbidden)
		return
	}

	l.Debug().Msg("Retrieved wishlist")
	json.NewEncoder(w).Encode(CreateSingleWishlistResponse(wishlist))
}

func (h *WishlistHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(internal.UserIDCtxKey).(string)
	params := mux.Vars(r)
	wishlistId := params["wishlist_id"]
	l := zerolog.Ctx(r.Context())

	var requestPayload Wishlist
	if err := json.NewDecoder(r.Body).Decode(&requestPayload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(wishlistId, 10, 64)
	if err != nil {
		l.Err(err).Msg("Failed to parse wishlist ID")
		http.Error(w, "Invalid wishlist ID", http.StatusBadRequest)
		return
	}

	requestPayload.ID = uint(id)
	requestPayload.OwnerId = userID

	updatedWishlist, err := h.wishlistService.Update(&requestPayload)
	if err != nil {
		l.Err(err).Msg("Failed to update wishlist")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	l.Debug().Int("WishlistID", int(updatedWishlist.ID)).Msg("Updated wishlist")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CreateSingleWishlistResponse(updatedWishlist))
}

func (h *WishlistHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(internal.UserIDCtxKey).(string)
	params := mux.Vars(r)
	wishlistId := params["wishlist_id"]
	l := zerolog.Ctx(r.Context())

	wishlistToDel, err := h.wishlistService.GetByID(wishlistId)
	if err != nil {
		l.Err(err).Msg("Failed to retrieve wishlist for deletion")
		http.Error(w, "Wishlist not found", http.StatusNotFound)
		return
	}

	if wishlistToDel.OwnerId != userID {
		l.Warn().Msg("User attempted to delete a wishlist they do not own")
		http.Error(w, "Forbidden: You are not the owner of this wishlist", http.StatusForbidden)
		return
	}

	if err := h.wishlistService.Delete(wishlistId); err != nil {
		l.Err(err).Msg("Failed to delete wishlist")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	l.Debug().Int("WishlistID", int(wishlistToDel.ID)).Msg("Deleted wishlist")
	w.WriteHeader(http.StatusNoContent)
}

func (h *WishlistHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", h.GetAll).Methods("GET")
	r.HandleFunc("/", h.Create).Methods("POST")
	r.HandleFunc("/{id}", h.GetByID).Methods("GET")
	r.HandleFunc("/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/{id}", h.DeleteByID).Methods("DELETE")
}
