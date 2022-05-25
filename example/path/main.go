package main

import (
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func Equal(ctx *fasthttp.RequestCtx) {
	_, _ = ctx.WriteString(`Hello World`)
}

func Contains(ctx *fasthttp.RequestCtx) {
	_, _ = ctx.WriteString(`24sdfq23rwasdfasdfHelloadfasdf23sadfasdfef2`)
}

func Pattern(ctx *fasthttp.RequestCtx) {
	_, _ = ctx.WriteString(`{"code": 200, "msg": "ok", "data":{}}`)
}

func loadRouter() *fasthttprouter.Router {
	r := fasthttprouter.New()
	r.GET("/equal", Equal)
	r.POST("/contains", Contains)
	r.GET("/pattern", Pattern)

	return r
}

func main() {
	r := loadRouter()

	server := fasthttp.Server{
		Handler: r.Handler,
	}

	err := server.ListenAndServe(":8080")
	if err != nil {
		log.Fatalf("server start error:%s\n", err.Error())
	}
}
