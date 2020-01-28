package entity

type Advert struct {
	Id                string
	Name              string
	MediaURL          string
	MediaKey          string
	ThumbnailURL      string
	MediaType         string
	ConversionCount   int
	ConversionMessage string
	ConversionType    string
	CampaignType      string
}

//advert_id = db.Column(db.String(256), index=True, primary_key=True, unique=True)
//name = db.Column(db.String(1024), index=True)
//media_url = db.Column(db.String(1024), index=True)
//media_key = db.Column(db.String(512), index=True)
//thumbnail_url = db.Column(db.String(1024), index=True)
//media_type = db.Column(db.String(10), index=True)
//conversion_count = db.Column(db.Integer, default=0)
//conversion_url = db.Column(db.String(1024), index=True)
//conversion_message = db.Column(db.String(128), index=True)
//conversion_type = db.Column(db.String(128), index=True)
//campaign_type = db.Column(db.String(24), default='advert', nullable=False)
