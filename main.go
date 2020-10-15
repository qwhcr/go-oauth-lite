package main

import (
	"log"

	"go-oauth-lite/handlers"
	oauth2handlers "go-oauth-lite/handlers/oauth2"
	configUtil "go-oauth-lite/util/config"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func main() {

	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	configUtil.ReadConfig()

	log.Println("starting covid tracker backend...")

	router := routing.New()

	router.Get("/health", handlers.HealthHandler)
	oauth2Routing := router.Group("/oauth2")
	oauth2Routing.Get("/authorize", oauth2handlers.AuthorizationHandler)

	log.Println("starting HTTP server on port " + configUtil.GetConfig().Port)
	fasthttp.ListenAndServe(":"+configUtil.GetConfig().Port, router.HandleRequest)

}
