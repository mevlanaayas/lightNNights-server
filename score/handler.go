package score

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"lightNNights/errors"
	"net/http"
)

type Handler struct {
	commandService CommandService
	queryService   QueryService
}

func NewHandler(commandService CommandService,
	queryService QueryService) *Handler {
	return &Handler{
		commandService: commandService,
		queryService:   queryService,
	}

}

func (receiver Handler) Save(ctx echo.Context) error {
	err := receiver.save(ctx)
	if err != nil {
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errorCode":    err.Code(),
			"errorMessage": err.Message(),
		})
	}
	response := BaseResponse{
		Success: true,
		Result:  SingleMessageResponse{Message: "saved"}}
	return ctx.JSON(http.StatusOK, response)
}

func (receiver Handler) save(ctx echo.Context) errors.Error {
	var request SaveScoreRequest
	err := ctx.Bind(&request)
	if err != nil {
		return errors.New("error while binding body to save score contract ", 1, fmt.Errorf("error while binding body to save score contract \n\t%v", err))
	}

	err = request.Validate()
	if err != nil {
		return errors.New(err.Error(), 1, fmt.Errorf("error while validating save score contract \n\t%v", err))
	}

	score := Score{
		Id:    request.Id,
		Point: request.Point,
	}
	err = receiver.commandService.Save(score)

	if err != nil {
		return errors.New("error while saving score", 1, fmt.Errorf("error while saving score \n\t%v", err))
	}
	return nil
}

func (receiver Handler) List(ctx echo.Context) error {
	scores, err := receiver.list()
	if err != nil {
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errorCode":    err.Code(),
			"errorMessage": err.Message(),
		})
	}
	response := BaseResponse{
		Success: true,
		Result:  scores,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (receiver Handler) list() ([]Score, errors.Error) {
	scores, err := receiver.queryService.List()

	if err != nil {
		return nil, errors.New("error while listing score", 1, fmt.Errorf("error while listing score \n\t%v", err))
	}
	return scores, nil
}
