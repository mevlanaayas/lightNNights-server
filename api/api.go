package api

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"lightNNights/config"
	"lightNNights/score"
	"net/http"
	"time"
)

type API struct {
	log          *logrus.Entry
	config       *config.Config
	echo         *echo.Echo
	scoreHandler *score.Handler
}

func NewAPI(config *config.Config, scoreHandler *score.Handler) *API {
	api := &API{
		config:       config,
		log:          logrus.StandardLogger().WithField("component", "score-api"),
		scoreHandler: scoreHandler,
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(api.configureRequestLogging)

	// Routes
	e.GET("/info", api.Info)

	v1 := e.Group("/v1")

	g := v1.Group("/score")
	g.GET("", nil)
	g.POST("", nil)

	api.echo = e
	return api
}

func (receiver *API) Start() error {
	return receiver.echo.Start(fmt.Sprintf(":%d", receiver.config.Port))
}

func (receiver *API) Info(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{
		"version":     "0.0.1",
		"description": "I am the best",
		"name":        "score api",
	})
}

func (receiver *API) configureRequestLogging(f echo.HandlerFunc) echo.HandlerFunc {

	return func(ctx echo.Context) error {
		request := ctx.Request()

		logger := receiver.log.WithFields(logrus.Fields{
			"method":         request.Method,
			"path":           request.URL.Path,
			"request_id":     uuid.NewString(),
			"user_agent":     request.UserAgent(),
			"content_length": request.ContentLength,
		})
		ctx.Set("logger", logger)
		startTime := time.Now()

		defer func() {
			response := ctx.Response()
			logger.WithFields(logrus.Fields{
				"status_code":  response.Status,
				"runtime_nano": time.Since(startTime).Nanoseconds(),
			}).Info("Finished request")
		}()

		logger.Info("Starting request")

		err := f(ctx)
		if err != nil {
			ctx.Error(err)
		}

		return err
	}
}
