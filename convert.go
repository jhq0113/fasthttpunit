package fasthttpunit

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

func CaseToRequest(api *Api, c Case) *fasthttp.Request {
	req := fasthttp.AcquireRequest()

	req.Header.SetMethod(api.Method)

	uri := api.Path
	if len(c.Params) > 0 {
		switch api.Method {
		case http.MethodGet, http.MethodHead, http.MethodDelete:
			uri += "?" + c.Params
		case http.MethodPost, http.MethodPut:
			req.SetBodyString(c.Params)
		}
	}

	req.SetRequestURI(uri)

	if len(c.Header) > 0 {
		for key, value := range api.Header {
			req.Header.Set(key, value)
		}
	}

	if len(c.Header) > 0 {
		for key, value := range c.Header {
			req.Header.Set(key, value)
		}
	}
	req.Header.SetHost(api.Host)

	if api.ContentType != "" {
		req.Header.SetContentType(api.ContentType)
	}

	return req
}
