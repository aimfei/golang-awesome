package base

type ResultData struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) *ResultData {
	return &ResultData{"00000", "success", data}
}
