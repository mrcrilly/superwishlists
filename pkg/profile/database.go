package profile

import (
	"errors"
)

// DBI is the database interface for the Profile model.
// It's designed to abstract the database from the mode
// so that the underlaying data store can be swapped out.
type DBI interface {
	GetByID(string) (Profile, error)
	GetByRealName(string) (Profile, error)
	GetByEmail(string) (Profile, error)

	Add(Profile) error
	Delete(string) error
	Update(Profile) error
}

// DB is a concrete implementation for he DBI
// interface. It will be used by HTTPS handlers for manipulating
// Profile objects in the underlaying database engine.
type DB struct {
	// STUB
}

// GetByID will fetch a Profile object based on the ID field.
func (d DB) GetByID(string) (Profile, error) { return Profile{}, errors.New("stub") }

// GetByRealName will fetch a Profile object based on the RealName field.
func (d DB) GetByRealName(string) (Profile, error) { return Profile{}, errors.New("stub") }

// GetByEmail will fetch a Profile based on the email address.
func (d DB) GetByEmail(string) (Profile, error) { return Profile{}, errors.New("stub") }

// Add will create a new Profile in the database based on the fields provided.
func (d DB) Add(Profile) error { return errors.New("stub") }

// Delete will use the provided ID to remove a Profile from the database.
func (d DB) Delete(string) error { return errors.New("stub") }

// Update will perform an "upsert" into the database, updating the profile
// if it exists or adding a whole new reocrd if it does not.
func (d DB) Update(Profile) error { return errors.New("stub") }
