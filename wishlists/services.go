package wishlists

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

func (s *WishlistService) GetAllByUserID(userID string) ([]*Wishlist, error) {
	return s.WishlistRepository.GetAllByUserID(userID)
}

func (s *WishlistService) Update(wishlist *Wishlist) (*Wishlist, error) {
	return s.WishlistRepository.Update(wishlist)
}

func (s *WishlistService) Delete(id string) error {
	return s.WishlistRepository.Delete(id)
}
