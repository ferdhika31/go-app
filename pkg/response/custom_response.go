package response

import (
	"encoding/json"
	"github.com/ferdhika31/go-app/pkg/logger"
	"github.com/labstack/echo/v4"
)

type CustomResponse interface {
	Json(e echo.Context, appResponseCode string, data interface{})
	JsonWithPagination(e echo.Context, appResponseCode string, data interface{}, meta ResponseMeta)
	JsonWithMessage(e echo.Context, appResponseCode string, data interface{}, message string)
	JsonWithCaptureError(e echo.Context, err error)
}

// Json is
func Json(e echo.Context, appResponseCode string, data interface{}) error {
	appResponse := GetAppResponse(appResponseCode)

	return e.JSON(appResponse.HttpStatus, Response{
		Code:    appResponse.Code,
		Message: appResponse.Messages,
		Data:    data,
	})
}

func JsonWithPagination(e echo.Context, appResponseCode string, data interface{}, meta ResponseMeta) error {
	appResponse := GetAppResponse(appResponseCode)

	return e.JSON(appResponse.HttpStatus, ResponsePagination{
		Code:    appResponse.Code,
		Message: appResponse.Messages,
		Data:    data,
		Meta:    meta,
	})
}

func JsonWithMessage(e echo.Context, appResponseCode string, data interface{}, message string) error {
	appResponse := GetAppResponse(appResponseCode)

	if message == "" {
		message = appResponse.Messages
	}

	return e.JSON(appResponse.HttpStatus, Response{
		Code:    appResponse.Code,
		Message: message,
		Data:    data,
	})
}

func JsonWithCaptureError(e echo.Context, err error) error {
	var pureError error = err
	var capt CaptureError
	err = json.Unmarshal([]byte(err.Error()), &capt)
	logger.Infof("capt: %v", capt)

	appResponse := GetAppResponse(capt.Code)
	messages := capt.Message
	if messages == "" {
		messages = appResponse.Messages
	}

	logger.Infof("Error: %v", err)

	if err != nil {
		return e.JSON(appResponse.HttpStatus, Response{
			Code:    appResponse.Code,
			Message: pureError.Error(),
		})
	} else {
		if capt.Trace {
			return e.JSON(appResponse.HttpStatus, Response{
				Code:    InternalServerError,
				Errors:  capt.Errors,
				Message: appResponse.Messages,
			})
		} else {
			return e.JSON(appResponse.HttpStatus, Response{
				Code:    appResponse.Code,
				Message: messages,
			})
		}
	}
}
