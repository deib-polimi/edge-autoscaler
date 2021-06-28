package slpaclient

// ResponseSLPA wraps the communities generated by SLPA
type ResponseSLPA struct {
	Communities []Community `json:"communities"`
}

// Community contains the community leader and its members
type Community struct {
	Name    string `json:"name"`
	Members []Host `json:"members"`
}