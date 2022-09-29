package http

type optionSuccess func(*Success)
type optionFailed func(*Failed)

func NewResponseDesc(en string) ResponseDesc {
	return ResponseDesc{EN: en}
}

func NewMeta(version string, status string, apiEnvironment string) Meta {
	return Meta{Version: version, Status: status, ApiEnvironment: apiEnvironment}
}

func NewSuccess(responseCode string, responseDesc ResponseDesc, data interface{}, httpStatus int, meta Meta) *Success {
	return &Success{
		ResponseCode: responseCode,
		ResponseDesc: responseDesc,
		Data:         data,
		HttpStatus:   httpStatus,
		Meta:         meta,
	}
}

func NewFailed(responseCode string, responseDesc ResponseDesc, httpStatus int, meta Meta) *Failed {
	return &Failed{
		ResponseCode: responseCode,
		ResponseDesc: responseDesc,
		HttpStatus:   httpStatus,
		Meta:         meta,
	}
}
