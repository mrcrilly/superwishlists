package friend

// Status represents whether or not a friend request
// has been accepted, rejected, or has no response yet.
type Status uint

const (
	noresponse = Status(0)
	accepted   = Status(1)
	rejected   = Status(2)
)

// Friend represents a user's friend and links them
// to their account, making them available for wish list
// invitations.
type Friend struct {
	ID              string
	OwnerID         string
	FriendProfileID string
	Status          Status
}
