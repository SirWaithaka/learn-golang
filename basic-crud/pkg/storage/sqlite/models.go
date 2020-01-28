package sqlite

import "time"

// Device model
type Device struct {
	UUID        string
	AuthStatus  bool
	Token       string
	PlaceID     string
	PlaceName   string
	LastUpdated time.Time
	Active      bool
}
