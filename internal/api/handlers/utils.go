package handlers

//ErrResponse api error response common structure
type ErrResponse struct {
	Errors      []string `json:"errors"`
	HumanErrors []string `json:"human_errors"`
}
