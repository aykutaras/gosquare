package base

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
