package login

import "errors"

// DBI is the database interface for the Login model.
// It's designed to abstract the database from the mode
// so that the underlaying data store can be swapped out.
type DBI interface {
	GetByID(string) (Login, error)
	GetByOwnerID(string) (Login, error)
	GetByUsername(string) (Login, error)

	Add(Login) error
	Delete(string) error
	Update(Login) error
}

// DB is a concrete implementation for he DBI
// interface. It will be used by HTTPS handlers for manipulating
// Login objects in the underlaying database engine.
type DB struct {
	// STUB
}

// GetByID will fetch a Login object from the database
// using the provided ID.
func (l DB) GetByID(string) (Login, error) { return Login{}, errors.New("stub") }

// GetByOwnerID will fetch a Login object from the database
// using the owner ID (the Profile that owns the login object.)
func (l DB) GetByOwnerID(string) (Login, error) { return Login{}, errors.New("stub") }

// GetByUsername will fetch a Login object from the database
// using the username of the Login object.
func (l DB) GetByUsername(string) (Login, error) { return Login{}, errors.New("stub") }

// Add will allow us to add a new Login to the underlaying database
// using the provided Login object and its fields.
func (l DB) Add(Login) error { return errors.New("stub") }

// Delete will let us delete a Login object from the database.
func (l DB) Delete(string) error { return errors.New("stub") }

// Update is used to "upsert" a record into the database using the
// provided Login object and its fields. Any existing fields are
// updated.
func (l DB) Update(Login) error { return errors.New("stub") }
