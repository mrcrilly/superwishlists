package wishlistitem

import "errors"

// DBI is the database interface for the WishListItem model.
// It's designed to abstract the database from the mode
// so that the underlaying data store can be swapped out.
type DBI interface {
	GetByID(string) (WishListItem, error)
	GetByOwnerID(string) (WishListItem, error)
	GetByListID(string) (WishListItem, error)

	Add(WishListItem) error
	Delete(string) error
	Update(WishListItem) error
}

// DB is a concrete implementation for he DBI
// interface. It will be used by HTTPS handlers for manipulating
// WishListItem objects in the underlaying database engine.
type DB struct {
	// STUB
}

// GetByID will fetch a WishListItem object based on the ID field.
func (d DB) GetByID(string) (WishListItem, error) { return WishListItem{}, errors.New("stub") }

// GetByOwnerID will fetch a WishListItem object based on the ID field.
func (d DB) GetByOwnerID(string) (WishListItem, error) { return WishListItem{}, errors.New("stub") }

// GetByListID will fetch a WishListItem object based on the ID field.
func (d DB) GetByListID(string) (WishListItem, error) { return WishListItem{}, errors.New("stub") }

// Add will create a new WishListItem in the database based on the fields provided.
func (d DB) Add(WishListItem) error { return errors.New("stub") }

// Delete will use the provided ID to remove a WishListItem from the database.
func (d DB) Delete(string) error { return errors.New("stub") }

// Update will perform an "upsert" into the database, updating the wish list
// if it exists or adding a whole new reocrd if it does not.
func (d DB) Update(WishListItem) error { return errors.New("stub") }
