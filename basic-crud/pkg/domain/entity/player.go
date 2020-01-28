package entity

import "time"

type Player struct {
	Id          string
	AuthStatus  bool
	Token       string
	PlaceId     string
	PlaceName   string
	LastUpdated time.Time
	Active      bool
}

//player_id = db.Column(db.String(256), index=True, primary_key=True)
//auth_status = db.Column(db.Boolean, default=False, nullable=False)
//token = db.Column(db.String(1024))
//place_id = db.Column(db.String(256), default='', nullable=False)
//place_name = db.Column(db.String(1024), default='', nullable=False)
//last_updated = db.Column(db.String(128), default='', nullable=False)
//active = db.Column(db.Boolean, default=False, nullable=False)
//state = db.Column(db.String(128), default='', nullable=False)
