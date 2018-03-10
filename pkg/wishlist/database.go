package wishlist

import (
	"errors"
)

// DBI is the database interface for the WishList model.
// It's designed to abstract the database from the mode
// so that the underlaying data store can be swapped out.
type DBI interface {
	GetByID(string) (WishList, error)
	GetByOwnerID(string) (WishList, error)

	Add(WishList) error
	Delete(string) error
	Update(WishList) error

	// GetItems(WishList) ([]wishlistitem.WishListItem, error)
	// GetFriends(WishList) ([]friend.Friend, error)
	// InviteFriends(string, []friend.Friend) error
	// RemoveFriends(string, []friend.Friend) error
}

// DB is a concrete implementation for he DBI
// interface. It will be used by HTTPS handlers for manipulating
// WishList objects in the underlaying database engine.
type DB struct {
	// STUB
}

// GetByID will fetch a WishList object based on the ID field.
func (d DB) GetByID(string) (WishList, error) { return WishList{}, errors.New("stub") }

// GetByOwnerID will fetch a WishList object based on the ID field.
func (d DB) GetByOwnerID(string) (WishList, error) { return WishList{}, errors.New("stub") }

// Add will create a new WishList in the database based on the fields provided.
func (d DB) Add(WishList) error { return errors.New("stub") }

// Delete will use the provided ID to remove a WishList from the database.
func (d DB) Delete(string) error { return errors.New("stub") }

// Update will perform an "upsert" into the database, updating the wish list
// if it exists or adding a whole new reocrd if it does not.
func (d DB) Update(WishList) error { return errors.New("stub") }

// func (d DB) GetItems(WishList) ([]wishlistitem.WishListItem, error) {
// 	return []wishlistitem.WishListItem{}, errors.New("stub")
// }
// func (d DB) GetFriends(WishList) ([]friend.Friend, error) {
// 	return []friend.Friend{}, errors.New("stub")
// }
// func (d DB) InviteFriends(string, []friend.Friend) error { return errors.New("stub") }
// func (d DB) RemoveFriends(string, []friend.Friend) error { return errors.New("stub") }
