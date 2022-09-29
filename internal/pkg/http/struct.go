package http

type ResponseDesc struct {
	EN string `json:"en"`
}

type Meta struct {
	Version        string `json:"version"`
	Status         string `json:"status"`
	ApiEnvironment string `json:"api_environment"`
}

type Success struct {
	ResponseCode string       `json:"response_code"`
	ResponseDesc ResponseDesc `json:"response_desc"`
	Data         interface{}  `json:"data"`
	HttpStatus   int          `json:"http_status"`
	Meta         Meta         `json:"meta"`
}

type Failed struct {
	ResponseCode string       `json:"response_code"`
	ResponseDesc ResponseDesc `json:"response_desc"`
	HttpStatus   int          `json:"http_status"`
	Meta         Meta         `json:"meta"`
}
