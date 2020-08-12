package entity

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type EventType string

const (
	EventTypeAction      EventType = "action"
	EventTypePageview    EventType = "pageview"
	EventTypeScreenview  EventType = "screenview"
	EventTypeSession     EventType = "session"
	EventTypeGroup       EventType = "group"
	EventTypeTransaction EventType = "transaction"
	EventTypeTimed       EventType = "timed_event"
)

func NewPageview() *RawMessage {
	return &RawMessage{
		EventID:   uuid.New().String(),
		Event:     string(EventTypePageview),
		Timestamp: time.Now().UTC().String(),
	}
}

func RawMessageFromPayload(msg *message.Message) (*RawMessage, error) {
	if msg.Payload == nil {
		return nil, errors.New("payload is nil")
	}

	var rm RawMessage
	if err := json.Unmarshal(msg.Payload, &rm); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal payload to raw_message")
	}

	return &rm, nil
}

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
	OS         string `json:"os" query:"os"`
	OSVersion  string `json:"osv" query:"osv"`
	Resolution string `json:"res" query:"res"`
	Viewport   string `json:"vp" query:"vp"`
	ColorDepth string `json:"cd" query:"cd"`
	Timezone   string `json:"tz" query:"tz"`
	Language   string `json:"lang" query:"lang"`
	IPAddress  string `json:"ip" query:"ip"`
	UserAgent  string `json:"ua" query:"ua"`
	IsMobile   string `json:"mob" query:"mob"`
	IsDesktop  string `json:"dsk" query:"dsk"`
	IsTablet   string `json:"tab" query:"tab"`
	IsBot      string `json:"bot" query:"bot"`

	//Browser
	Browser        string `json:"b" query:"b"`
	BrowserName    string `json:"bn" query:"bn"`
	BrowserVersion string `json:"bv" query:"bv"`
	BrowserEngine  string `json:"be" query:"be"`

	//Action
	Category string `json:"cat" query:"cat"`
	Action   string `json:"act" query:"act"`
	Label    string `json:"lab" query:"lab"`
	Property string ` json:"prop" query:"prop"`
	Value    string `json:"val" query:"val"`

	//Pageview
	URL      string `json:"url" query:"url"`
	Path     string `json:"path" query:"path"`
	Title    string `json:"page" query:"page"`
	Referrer string `json:"refr" query:"refr"`

	//Screenview
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

func (rm *RawMessage) GetApp(ar AppReporter) (*App, error) {
	aid, err := strconv.Atoi(rm.AppID)
	if err == nil {
		a, err := ar.Get(aid)
		if err == nil {
			return a, nil
		}
	}

	return ar.GetByTrackingID(rm.TrackingID)
}

func (rm *RawMessage) GetSession(sr SessionReporter) (*Session, error) {
	sid, err := uuid.Parse(rm.SessionID)
	if err != nil {
		return nil, err
	}

	return sr.Get(sid)
}

func (rm *RawMessage) GetUser(ur UserReporter) (*User, error) {
	return ur.Get(rm.UserID)
}
