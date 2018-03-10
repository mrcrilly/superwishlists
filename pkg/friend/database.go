package friend

import "errors"

// DBI is the database interface for the Friend model.
// It's designed to abstract the database from the mode
// so that the underlaying data store can be swapped out.
type DBI interface {
	GetByID(string) (Friend, error)
	GetByOwnerID(string) (Friend, error)
	GetByProfileID(string) (Friend, error)

	Add(Friend) error
	Delete(string) error
	Update(Friend) error
}

// DB is a concrete implementation for he DBI
// interface. It will be used by HTTPS handlers for manipulating
// Friend objects in the underlaying database engine.
type DB struct {
	// STUB
}

// GetByID will fetch a Friend object based on the ID field.
func (d DB) GetByID(string) (Friend, error) { return Friend{}, errors.New("stub") }

// GetByOwnerID will fetch a Friend object based on the ID field.
func (d DB) GetByOwnerID(string) (Friend, error) { return Friend{}, errors.New("stub") }

// GetByProfileID will fetch a Friend object based on the ID field.
func (d DB) GetByProfileID(string) (Friend, error) { return Friend{}, errors.New("stub") }

// Add will create a new Friend in the database based on the fields provided.
func (d DB) Add(Friend) error { return errors.New("stub") }

// Delete will use the provided ID to remove a Friend from the database.
func (d DB) Delete(string) error { return errors.New("stub") }

// Update will perform an "upsert" into the database, updating the wish list
// if it exists or adding a whole new reocrd if it does not.
func (d DB) Update(Friend) error { return errors.New("stub") }
