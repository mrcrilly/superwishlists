package profile

import (
	"time"

	"github.com/mrcrilly/superwishlists/pkg/login"
)

type Profile struct {
	ID       string
	Logins   []login.Login
	RealName string
	Email    string
	DOB      time.Time
	// Friends []friend.Friend
	// WishLists []wishlist.WishList
}
