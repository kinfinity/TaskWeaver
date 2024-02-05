package rest_client

import (
	"encoding/json"
)

// An interface for response objects.
type IResponse interface {
	Normalize()
	onError(ErrorHandler func(...any))
}

// Base Response across rest calls
type clientResponse struct {
	Error   RestError
	Payload any
}

// Normalize is used to normalize the response data.
func (cr *clientResponse) Normalize() {
	// Unmarshal the response JSON directly into the existing instance
	json_data, _err := json.Marshal(cr.Payload)
	if _err != nil {
		if err := json.Unmarshal(json_data, cr); err != nil {
			// Handle unmarshal error
			return
		}

	}
}
