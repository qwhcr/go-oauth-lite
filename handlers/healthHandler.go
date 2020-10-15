package handlers

import (
	"fmt"
	"net/http"

	constants "go-oauth-lite/constants/handler"

	routing "github.com/qiangxue/fasthttp-routing"
)

func HealthHandler(c *routing.Context) error {
	_, err := fmt.Fprint(c.Response.BodyWriter(), constants.HealthResponse)
	if err != nil {
		return err
	}
	c.Response.SetStatusCode(http.StatusOK)
	return nil
}
