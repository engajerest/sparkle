package subscription
import "net/http"

func NewBadRequestErr(message string) *RestError {
	return &RestError{
		Status:  http.StatusBadRequest,
		Message: message,
		Error:   "bad_request",
	}
}