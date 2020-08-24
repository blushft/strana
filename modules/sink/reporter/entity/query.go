package entity

type QueryParams struct {
	AppIDs     []int             `json:"aid" query:"aid"`
	SessionIDs []string          `json:"sid" query:"sid"`
	UserIDs    []string          `json:"uid" query:"uid"`
	GroupIDs   []string          `json:"gid" query:"gid"`
	Offset     int               `json:"offset" query:"offset"`
	Limit      int               `json:"limit" query:"limit"`
	Start      int               `json:"start" query:"start"`
	End        int               `json:"end" query:"end"`
	Params     map[string]string `json:"params" query:"params"`
}
