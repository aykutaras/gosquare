package users

type UserCheckIns struct {
	Meta          Meta          `json:"meta"`
	Notifications Notifications `json:"notifications"`
	Response      Response      `json:"response"`
}

type Meta struct {
	Code         int    `json:"code"`
	ErrorType    string `json:"errorType"`
	ErrorDetails string `json:"errorDetails"`
}

type Notifications []Notification

type Notification struct {
	Type string      `json:"type"`
	Item interface{} `json:"item"`
}

type Response struct {
	CheckIns CheckIns `json:"checkins"`
}

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
	Likes            Likes    `json:"likes"`
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
	Gender       string    `json:"gender"`
	Relationship string    `json:"relationship"`
	Photo        UserPhoto `json:"photo"`
}

type Venue interface {
}

type Likes struct {
	Count  int         `json:"count"`
	Groups []UserGroup `json:"groups"`
}

type UserGroup struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
	Items []User `json:"items"`
}

type UserPhoto struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

type Sticker interface {
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
