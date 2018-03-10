package wishlistitem

// WishListItem is the model for an item added to
// a WishList.
type WishListItem struct {
	ID          string
	OwnerID     string
	ListID      string
	Title       string
	Description string
	Cost        float32
	Rating      uint
	Bought      bool
	BeingBought bool
	Active      bool
}
