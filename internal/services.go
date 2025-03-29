package internal

// Service layer
type WishlistService struct {
	WishlistRepository *WishlistRepository
}

func NewWishlistService(wishlistRepository *WishlistRepository) *WishlistService {
	return &WishlistService{
		WishlistRepository: wishlistRepository,
	}
}

func (s *WishlistService) Create(wishlist *Wishlist) error {
	return s.WishlistRepository.Create(wishlist)
}

func (s *WishlistService) GetByID(id string) (*Wishlist, error) {
	return s.WishlistRepository.GetByID(id)
}
