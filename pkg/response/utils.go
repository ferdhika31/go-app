package response

import (
	"encoding/json"
	"fmt"
)

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}

func GetAppResponse(appResponseCode string) AppResponse {
	return appResponses[appResponseCode]
}

func CustomException(code string, message string) error {
	err, _ := json.Marshal(CaptureError{Trace: false, Code: code, Message: message})
	return fmt.Errorf(string(err))
}

func CustomExceptionWithTrace(code string, errs string) error {
	err, _ := json.Marshal(CaptureError{Trace: true, Code: code, Message: errs, Errors: errs})
	return fmt.Errorf(string(err))
}
