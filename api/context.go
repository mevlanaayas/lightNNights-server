package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func getLogger(ctx echo.Context) *logrus.Entry {
	obj := ctx.Get("logger")
	if obj == nil {
		return logrus.NewEntry(logrus.StandardLogger())
	}

	return obj.(*logrus.Entry)
}
