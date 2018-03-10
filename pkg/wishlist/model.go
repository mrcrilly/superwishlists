package wishlist

// WishList represents the Wish List model.
type WishList struct {
	ID          string
	OwnerID     string
	Title       string
	Description string
	Active      bool
	Private     bool
	// Items []wishlistitem.WishListItem
	// Friends []friend.Friend
}
