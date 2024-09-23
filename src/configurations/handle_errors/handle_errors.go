package handle_errors

import (
	"net/http"
)

// it needs to have `json: field_name_in_json` to convert to json
type GeneralError struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Reasons []Reason `json:"reasons"`
}

type Reason struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func FormatGeneralError(message, err string, code int, reasons []Reason) *GeneralError {
	return &GeneralError{
		Message: message,
		Err:     err,
		Code:    code,
		Reasons: reasons,
	}
}

func (r *GeneralError) Error() string {
	return r.Message
}

func GenerateBadRequestErr(message string) *GeneralError {
	return &GeneralError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func GenerateBadRequestValidationtErr(message string, reasons []Reason) *GeneralError {
	return &GeneralError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Reasons: reasons,
	}
}

func InternalServerErr(message string) *GeneralError {
	return &GeneralError{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func GenerateNotFoundErr(message string) *GeneralError {
	return &GeneralError{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func GenerateForbiddenError(message string) *GeneralError {
	return &GeneralError{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
