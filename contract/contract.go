package contract

type (
	AdapterResult struct {
		Success bool        `json:"success"`
		Result  interface{} `json:"result"`
	}
	BaseResponse struct {
		Success bool        `json:"success"`
		Result  interface{} `json:"result"`
	}
)
