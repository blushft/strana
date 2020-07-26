package entity

import "encoding/json"

type EventType string

const (
	EventTypeAction     EventType = "action"
	EventTypePageview             = "pageview"
	EventTypeScreenview           = "screenview"
	EventTypeSession              = "session"
	EventTypeGroup                = "group"
)

type RawMessage struct {
	AppID      string `json:"aid" query:"aid"`
	SessionID  string `json:"sid" query:"sid"`
	UserID     string `json:"uid" query:"uid"`
	GroupID    string `json:"gid" query:"gid"`
	DeviceID   string `json:"cid" query:"cid"`
	TrackingID string `json:"tid" query:"tid"`
	EventID    string `json:"eid" query:"eid"`

	NewSession string `json:"news" query:"news"`
	NewUser    string `json:"newu" query:"newu"`
	NewDevice  string `json:"newd" query:"newd"`

	// Common
	Event          string `json:"e" query:"e"`
	NonInteractive string `json:"ni" query:"ni"`
	DataSource     string `json:"ds" query:"ds"`
	Namespace      string `json:"ns" query:"ns"`
	Platform       string `json:"p" query:"p"`
	AppVersion     string `json:"av" query:"av"`
	Timestamp      string `json:"dtm" query:"dtm"`
	SentTimestamp  string `json:"stm" query:"stm"`
	TrueTimestamp  string `json:"ttm" query:"ttm"`

	//Device
	Resolution string `json:"res" query:"res"`
	Viewport   string `json:"vp" query:"vp"`
	ColorDepth string `json:"cd" query:"cd"`
	Timezone   string `json:"tz" query:"tz"`
	Language   string `json:"lang" query:"lang"`
	IPAddress  string `json:"ip" query:"ip"`
	UserAgent  string `json:"ua" query:"ua"`

	//Structured Event
	Category string `json:"cat" query:"cat"`
	Action   string `json:"act" query:"act"`
	Label    string `json:"lab" query:"lab"`
	Property string ` json:"prop" query:"prop"`
	Value    string `json:"val" query:"val"`

	// Pageview
	URL      string `json:"url" query:"url"`
	Path     string `json:"path" query:"path"`
	Title    string `json:"page" query:"page"`
	Referrer string `json:"refr" query:"refr"`

	// Screenview
	ScreenName string `json:"sn" query:"sn"`
	ScreenID   string `json:"scid" query:"scid"`

	// Timing
	TimingCategory string `json:"utc" query:"utc"`
	TimingVariable string `json:"utv" query:"utv"`
	TimingLabel    string `json:"utl" query:"utl"`
	TimingTime     string `json:"utt" query:"utt"`

	// Location
	GeoID      string `json:"geoid" query:"geoid"`
	Region     string `json:"reg" query:"reg"`
	Locale     string `json:"loc" query:"loc"`
	Country    string `json:"cnty" query:"cnty"`
	City       string `json:"city" query:"city"`
	PostalCode string `json:"zip" query:"zip"`
	Longitude  string `json:"long" query:"long"`
	Latitude   string `json:"lat" query:"lat"`

	Extra map[string]string `json:"ex" query:"ex"`
}

func (rm *RawMessage) JSON() ([]byte, error) {
	return json.Marshal(rm)
}
