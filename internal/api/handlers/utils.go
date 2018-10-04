package handlers

//handlers constant variables
var (
	SessionCookieName           = "session_token"
	SessionCookieExpirationTime = int64(120) //in seconds
)

//ErrResponse api error response common structure
type ErrResponse struct {
	Errors      []string `json:"errors"`
	HumanErrors []string `json:"human_errors"`
}
