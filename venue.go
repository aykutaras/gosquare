package gosquare

type Venue struct {
	Id         string          `json:"id"`
	Name       string          `json:"name"`
	Contact    Contact         `json:"contact"`
	Location   VenueLocation   `json:"location"`
	Categories []VenueCategory `json:"categories`
	Verified   bool            `json:"verified"`
	Stats      VenueStats      `json:"stats"`
	Url        string          `json:"url"`
	Like       bool            `json:"like"`
	VenuePage  VenuePage       `json:"venuePage"`
	StoreId    string          `json:"storeId"`
	Menu       VenueMenu       `json:"menu"`
}

type VenueLocation struct {
	Address          string   `json:"address"`
	CrossStreet      string   `json:"crossStreet"`
	Lat              float64  `json:"lat"`
	Lon              float64  `json:"lon"`
	PostalCode       string   `json:"postalCode"`
	CC               string   `json:"cc"`
	City             string   `json:"city"`
	State            string   `json:"state"`
	Country          string   `json:"country"`
	FormattedAddress []string `json:"formattedAddress"`
}

type VenueCategory struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	PluralName string `json:"pluralName"`
	ShortName  string `json:"shortName"`
	Icon       Icon   `json:"icon"`
	Primary    bool   `json:"primary"`
}

type VenueStats struct {
	CheckinsCount int `json:"checkinsCount"`
	UsersCount    int `json:"usersCount"`
	TipCount      int `json:"tipCount"`
}

type VenuePage struct {
	Id string `json:"id"`
}

type VenueMenu struct {
	Type        string `json:"type"`
	Label       string `json:"label"`
	Anchor      string `json:"anchor"`
	Url         string `json:"url"`
	MobileUrl   string `json:"mobileUrl"`
	ExternalUrl string `json:"externalUrl"`
}
