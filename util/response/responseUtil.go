package response

import (
	"encoding/json"

	routing "github.com/qiangxue/fasthttp-routing"
)

type responseWrapper struct {
	Payload interface{} `json:"payload"`
}

type errorPayload struct {
	ErrorMessage string `json:"error_message"`
	RawError     string `json:"raw_error"`
}

// Respond writes the payload and sets status code to the response
func Respond(ctx *routing.Context, payload interface{}, statusCode int) error {
	ctx.SetStatusCode(statusCode)
	res := responseWrapper{Payload: payload}
	data, err := json.Marshal(res)
	if err != nil {
		return err
	}
	err = ctx.WriteData(data)
	if err != nil {
		return err
	}
	return nil
}

// RespondError notifies caller the error message and raw error, also sets status code
func RespondError(ctx *routing.Context, err error, errorMessage string, statusCode int) error {

	e := errorPayload{
		ErrorMessage: errorMessage,
	}
	if err != nil {
		e.RawError = err.Error()
	}

	return Respond(ctx, e, statusCode)
}
