package event

import (
	"encoding/json"
	"net"

	"github.com/mitchellh/mapstructure"
)

type Action struct {
	Category string      `json:"category,omitempty" structs:"category,omitempty" mapstructure:"category,omitempty"`
	Action   string      `json:"action,omitempty" structs:"action,omitempty" mapstructure:"action,omitempty"`
	Label    string      `json:"label,omitempty" structs:"label,omitempty" mapstructure:"label,omitempty"`
	Property string      `json:"property,omitempty" structs:"property,omitempty" mapstructure:"property,omitempty"`
	Value    interface{} `json:"value,omitempty" structs:"value,omitempty" mapstructure:"value,omitempty"`
}

func decodeAction(v interface{}) (Context, error) {
	var ctx Action
	if err := mapstructure.Decode(v, &ctx); err != nil {
		return nil, err
	}

	return &context{
		typ: ContextAction,
		v:   ctx,
	}, nil
}

type Alias struct {
	From string
	To   string
}

func NewAliasContext(from, to string) Context {
	return &context{
		typ: ContextAlias,
		v: Alias{
			From: from,
			To:   to,
		},
	}
}

type App struct {
	Name       string                 `json:"name,omitempty" structs:"name,omitempty" mapstructure:"name,omitempty"`
	Version    string                 `json:"version,omitempty" structs:"version,omitempty" mapstructure:"version,omitempty"`
	Build      string                 `json:"build,omitempty" structs:"build,omitempty" mapstructure:"build,omitempty"`
	Namespace  string                 `json:"namespace,omitempty" structs:"namespace,omitempty" mapstructure:"namespace,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty" structs:"properties,omitempty" mapstructure:"properties,omitempty"`
}

type Browser struct {
	Name      string `json:"name,omitempty" structs:"name,omitempty" mapstructure:"name,omitempty"`
	Version   string `json:"version,omitempty" structs:"version,omitempty" mapstructure:"version,omitempty"`
	UserAgent string `json:"userAgent,omitempty" structs:"userAgent,omitempty" mapstructure:"userAgent,omitempty"`
}

func NewBrowserContext(ctx *Browser) Context {
	return newContext(ContextBrowser, ctx)
}

type Campaign struct {
	Name    string `json:"name,omitempty" structs:"name,omitempty" mapstructure:"name,omitempty"`
	Source  string `json:"source,omitempty" structs:"source,omitempty" mapstructure:"source,omitempty"`
	Medium  string `json:"medium,omitempty" structs:"medium,omitempty" mapstructure:"medium,omitempty"`
	Term    string `json:"term,omitempty" structs:"term,omitempty" mapstructure:"term,omitempty"`
	Content string `json:"content,omitempty" structs:"content,omitempty" mapstructure:"content,omitempty"`
}

type Connectivity struct {
	Bluetooth bool   `json:"bluetooth,omitempty" structs:"bluetooth,omitempty" mapstructure:"bluetooth,omitempty"`
	Cellular  bool   `json:"cellular,omitempty" structs:"cellular,omitempty" mapstructure:"cellular,omitempty"`
	WIFI      bool   `json:"wifi,omitempty" structs:"wifi,omitempty" mapstructure:"wifi,omitempty"`
	Ethernet  bool   `json:"ethernet,omitempty" structs:"ethernet,omitempty" mapstructure:"ethernet,omitempty"`
	Carrier   string `json:"carrier,omitempty" structs:"carrier,omitempty" mapstructure:"carrier,omitempty"`
}

type Device struct {
	ID           string                 `json:"id,omitempty" structs:"id,omitempty" mapstructure:"id,omitempty"`
	Manufacturer string                 `json:"manufacturer,omitempty" structs:"manufacturer,omitempty" mapstructure:"manufacturer,omitempty"`
	Model        string                 `json:"model,omitempty" structs:"model,omitempty" mapstructure:"model,omitempty"`
	Name         string                 `json:"name,omitempty" structs:"name,omitempty" mapstructure:"name,omitempty"`
	Type         string                 `json:"type,omitempty" structs:"type,omitempty" mapstructure:"type,omitempty"`
	Version      string                 `json:"version,omitempty" structs:"version,omitempty" mapstructure:"version,omitempty"`
	Mobile       bool                   `json:"mobile,omitempty" structs:"mobile,omitempty" mapstructure:"mobile,omitempty"`
	Tablet       bool                   `json:"tablet,omitempty" structs:"tablet,omitempty" mapstructure:"tablet,omitempty"`
	Desktop      bool                   `json:"desktop,omitempty" structs:"desktop,omitempty" mapstructure:"desktop,omitempty"`
	Properties   map[string]interface{} `json:"properties,omitempty" structs:"properties,omitempty" mapstructure:"properties,omitempty"`
}

func NewDeviceContext(ctx *Device) Context {
	return newContext(ContextDevice, ctx)
}

type Extra map[string]interface{}

type Library struct {
	Name    string `json:"name,omitempty" structs:"name,omitempty" mapstructure:"name,omitempty"`
	Version string `json:"version,omitempty" structs:"version,omitempty" mapstructure:"version,omitempty"`
}

func NewLibraryContext(name, version string) Context {
	return newContext(ContextLibrary, &Library{
		Name:    name,
		Version: version,
	})
}

type Location struct {
	Street     string  `json:"street,omitempty" structs:"street,omitempty" mapstructure:"street,omitempty"`
	City       string  `json:"city,omitempty" structs:"city,omitempty" mapstructure:"city,omitempty"`
	State      string  `json:"state,omitempty" structs:"state,omitempty" mapstructure:"state,omitempty"`
	PostalCode string  `json:"postalCode,omitempty" structs:"postalCode,omitempty" mapstructure:"postalCode,omitempty"`
	Region     string  `json:"region,omitempty" structs:"region,omitempty" mapstructure:"region,omitempty"`
	Locale     string  `json:"locale,omitempty" structs:"locale,omitempty" mapstructure:"locale,omitempty"`
	Country    string  `json:"country,omitempty" structs:"country,omitempty" mapstructure:"country,omitempty"`
	Longitude  float64 `json:"longitude,omitempty" structs:"longitude,omitempty" mapstructure:"longitude,omitempty"`
	Latitude   float64 `json:"latitude,omitempty" structs:"latitude,omitempty" mapstructure:"latitude,omitempty"`
	Timezone   string  `json:"timezone,omitempty" structs:"timezone,omitempty" mapstructure:"timezone,omitempty"`
}

func NewLocationContext(loc *Location) Context {
	return newContext(ContextLocation, loc)
}

type Network struct {
	IP        net.IP `json:"ip,omitempty" structs:"ip,omitempty" mapstructure:"ip,omitempty"`
	UserAgent string `json:"userAgent,omitempty" structs:"userAgent,omitempty" mapstructure:"userAgent,omitempty"`
}

func NewNetworkContext(ip, ua string) Context {
	nip := net.ParseIP(ip)

	return newContext(ContextNetwork, &Network{
		IP:        nip,
		UserAgent: ua,
	})
}

func (ctx *Network) MarshalJSON() ([]byte, error) {
	ipt, err := ctx.IP.MarshalText()
	if err != nil {
		return nil, err
	}

	m := map[string]interface{}{
		"ip":        string(ipt),
		"userAgent": ctx.UserAgent,
	}

	return json.Marshal(m)
}

func (ctx *Network) UnmarshalJSON(b []byte) error {
	m := map[string]interface{}{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	ua, ok := m["userAgent"]
	if ok {
		ctx.UserAgent = ua.(string)
	}

	ipt, ok := m["ip"]
	if ok {
		ipb := ipt.(string)
		ip := net.IP{}
		if err := ip.UnmarshalText([]byte(ipb)); err != nil {
			return err
		}

		ctx.IP = ip
	}

	return nil
}

type OS struct {
	Name    string `json:"name,omitempty" structs:"name,omitempty" mapstructure:"name,omitempty"`
	Version string `json:"version,omitempty" structs:"version,omitempty" mapstructure:"version,omitempty"`
}

func NewOSContext(name, version string) Context {
	return newContext(ContextOS, &OS{
		Name:    name,
		Version: version,
	})
}

type Page struct {
	Hash     string `json:"hash,omitempty" structs:"hash,omitempty" mapstructure:"hash,omitempty"`
	Path     string `json:"path,omitempty" structs:"path,omitempty" mapstructure:"path,omitempty"`
	Referrer string `json:"referrer,omitempty" structs:"referrer,omitempty" mapstructure:"referrer,omitempty"`
	Search   string `json:"search,omitempty" structs:"search,omitempty" mapstructure:"search,omitempty"`
	Title    string `json:"title,omitempty" structs:"title,omitempty" mapstructure:"title,omitempty"`
	Hostname string `json:"hostname,omitempty" structs:"hostname,omitempty" mapstructure:"hostname,omitempty"`
}

type Referrer struct {
	Type     string `json:"type,omitempty" structs:"type,omitempty" mapstructure:"type,omitempty"`
	Name     string `json:"name,omitempty" structs:"name,omitempty" mapstructure:"name,omitempty"`
	Hostname string `json:"hostname,omitempty" structs:"hostname,omitempty" mapstructure:"hostname,omitempty"`
	Link     string `json:"link,omitempty" structs:"link,omitempty" mapstructure:"link,omitempty"`
}

type Session struct{}

type Timing struct {
	Category string  `json:"category,omitempty" structs:"category,omitempty" mapstructure:"category,omitempty"`
	Label    string  `json:"label,omitempty" structs:"label,omitempty" mapstructure:"label,omitempty"`
	Unit     string  `json:"unit,omitempty" structs:"unit,omitempty" mapstructure:"unit,omitempty"`
	Variable string  `json:"variable,omitempty" structs:"variable,omitempty" mapstructure:"variable,omitempty"`
	Value    float64 `json:"value,omitempty" structs:"value,omitempty" mapstructure:"value,omitempty"`
}

func NewTimingContext(cat, label, variable string, value float64) Context {
	return newContext(ContextTiming, &Timing{
		Category: cat,
		Label:    label,
		Variable: variable,
		Value:    value,
	})
}

type Traits struct {
	Age         int                    `json:"age,omitempty" structs:"age,omitempty" mapstructure:"age,omitempty"`
	Birthday    string                 `json:"birthday,omitempty" structs:"birthday,omitempty" mapstructure:"birthday,omitempty"`
	Description string                 `json:"description,omitempty" structs:"description,omitempty" mapstructure:"description,omitempty"`
	Email       string                 `json:"email,omitempty" structs:"email,omitempty" mapstructure:"email,omitempty"`
	FirstName   string                 `json:"firstName,omitempty" structs:"firstName,omitempty" mapstructure:"firstName,omitempty"`
	LastName    string                 `json:"lastName,omitempty" structs:"lastName,omitempty" mapstructure:"lastName,omitempty"`
	Gender      string                 `json:"gender,omitempty" structs:"gender,omitempty" mapstructure:"gender,omitempty"`
	Phone       string                 `json:"phone,omitempty" structs:"phone,omitempty" mapstructure:"phone,omitempty"`
	Title       string                 `json:"title,omitempty" structs:"title,omitempty" mapstructure:"title,omitempty"`
	Username    string                 `json:"username,omitempty" structs:"username,omitempty" mapstructure:"username,omitempty"`
	Website     string                 `json:"website,omitempty" structs:"website,omitempty" mapstructure:"website,omitempty"`
	Extra       map[string]interface{} `json:"extra,omitempty" structs:"extra,omitempty" mapstructure:"extra,omitempty"`
}

type User struct {
	UserID string `json:"userID,omitempty" structs:"userID,omitempty" mapstructure:"userID,omitempty"`
	AnonID string `json:"anonID,omitempty" structs:"anonID,omitempty" mapstructure:"anonID,omitempty"`
	Name   string `json:"name,omitempty" structs:"name,omitempty" mapstructure:"name,omitempty"`
	Traits Traits `json:"traits,omitempty" structs:"traits,omitempty" mapstructure:"traits,omitempty"`
}

func NewUserContext(u *User) Context {
	return newContext(ContextUser, u)
}

type Viewport struct {
	Density int `json:"density,omitempty" structs:"density,omitempty" mapstructure:"density,omitempty"`
	Width   int `json:"width,omitempty" structs:"width,omitempty" mapstructure:"width,omitempty"`
	Height  int `json:"height,omitempty" structs:"height,omitempty" mapstructure:"height,omitempty"`
}
