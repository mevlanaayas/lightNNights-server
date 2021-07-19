package score

type (
	Request interface {
		Validate() error
	}
	BaseResponse struct {
		Success bool        `json:"success"`
		Result  interface{} `json:"result"`
	}
	SingleMessageResponse struct {
		Message string `json:"message"`
	}
	SaveScoreRequest struct {
		Id    string `json:"id"`
		Point int64  `json:"point"`
		Key   string `json:"key"`
	}

	Score struct {
		Id    string `bson:"id,omitempty"`
		Point int64  `bson:"point,omitempty"`
	}
)

func (receiver SaveScoreRequest) Validate() error {
	return nil
}
