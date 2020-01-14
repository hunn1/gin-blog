package helpers

type ApiReturn struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	//RedirectUrl string      `json:"redirect_url,omitempty"`
}

func NewApiReturn(code int, msg string, data interface{}) *ApiReturn {
	return &ApiReturn{
		code,
		msg,
		data,
		//redirect_url,
	}
}
