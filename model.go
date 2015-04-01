package gosquare

type CheckIns struct {
	Count int           `json:"count"`
	Items []CheckInItem `json:"items"`
}

type CheckInItem struct {
	Id               string   `json:"id"`
	CreatedAt        int      `json:"createdAt"`
	Type             string   `json:"type"`
	Shout            string   `json:"shout"`
	Entities         []Entity `json:"entities"`
	TimeZoneOffset   int      `json:"timeZoneOffset"`
	DisplayGeo       Geo      `json:"displayGeo"`
	ExactContextLine string   `json:"exactContextLine"`
	With             []User   `json:"with"`
	Venue            Venue    `json:"venue"`
	Likes            Group    `json:"likes"`
	Like             bool     `json:"like"`
	Sticker          Sticker  `json:"sticker"`
	Photos           Photos   `json:"photos"`
	Posts            Posts    `json:"posts"`
	Comments         Comments `json:"comments"`
	Source           Source   `json:"source"`
}

type Entity struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Indices []int  `json:indices"`
}

type Geo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id           string    `json:"id"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Photo        Icon      `json:"photo"`
	Relationship string    `json:"relationship"`
	Friends      Group     `json:"friends"`
	Type         string    `json:"type"`
	HomeCity     string    `json:"homeCity"`
	Gender       string    `json:"gender"`
	Contact      Contact   `json:"contact"`
	Bio          string    `json:"bio"`
	Followers    Group     `json:"followers"`
	Following    Group     `json:"following"`
	Venue        VenuePage `json:"venue"`
}

type Contact struct {
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	FormattedPhone   string `json:"formattedPhone"`
	Twitter          string `json:"twitter"`
	Facebook         string `json:"facebook"`
	FacebookUsername string `json:"facebookUsername"`
	FacebookName     string `json:"facebookName"`
}

type Group struct {
	Count  int         `json:"count"`
	Groups []UserGroup `json:"groups"`
}

type UserGroup struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
	Items []User `json:"items"`
}

type Icon struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

type Sticker struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	Image          Image          `json:"image"`
	StickerType    string         `json:"stickerType"`
	PickerPosition PickerPosition `json:"pickerPosition"`
	TeaseText      string         `json:"teaserText"`
}

type Image struct {
	Prefix string `json:"prefix"`
	Sizes  [2]int `json:"sizes"`
	Name   string `json:"name"`
}

type PickerPosition struct {
	Page  int `json:"page"`
	Index int `json:"index"`
}

type Photos struct {
	Count int     `json:"count"`
	Items []Photo `json:"items"`
}

type Photo struct {
	Id         string `json:"id"`
	CreatedAt  int    `json:"createdAt"`
	Source     Source `json:"source"`
	Prefix     string `json:"prefix"`
	Suffix     string `json:"suffix"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	User       User   `json:"user"`
	Visibility string `json:"visibility"`
}

type Posts struct {
	Count     int `json:"count"`
	TextCount int `json:"textCount"`
}

type Comments struct {
	Count int `json:"count"`
}

type Source struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
