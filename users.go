package gosquare

type Users struct {
	CheckIns UserCheckIns `json:"checkIns"`
	Profile  Profile      `json:"profile"`
	Friends  Friends      `json:"friends"`
}
