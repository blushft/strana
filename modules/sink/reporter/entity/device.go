package entity

type Device struct {
	ID          string `json:"id"`
	Fingerprint string `json:"fingerprint"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	IPAddress   string `json:"ip_address"`
	Resolution  string `json:"resolution"`
	Viewport    string `json:"viewport"`
	ColorDepth  string `json:"color_depth"`
	Timezone    string `json:"timezone"`
	Language    string `json:"language"`
	UserAgent   string `json:"user_agent"`
}
