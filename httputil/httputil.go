package httputil

// HTTPError represents an error response
//
//swagger:ignore
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
