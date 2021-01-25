package utilserrors

import (
	"encoding/json"
	"net/http"
)

// HTTPError is...
type HTTPError struct {
	StatusCode   int
	ErrorMessage string
}

// SendHTTPError sends http error
func SendHTTPError(w http.ResponseWriter, statusCode int, errorMessage string) {

	w.WriteHeader(statusCode)
	response := map[string]interface{}{
		"error": errorMessage,
	}
	json.NewEncoder(w).Encode(response)

}
