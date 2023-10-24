package error

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Error struct {
	CustomCode int `json:"code,omitempty"`

	*echo.HTTPError
}

func (e *Error) Error() string {
	return fmt.Sprintf("CustomCode=%d, %s", e.CustomCode, e.HTTPError.Error())
}
