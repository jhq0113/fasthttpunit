package fasthttpunit

import (
	"regexp"
	"strings"

	"github.com/valyala/fasthttp"
)

const (
	Equal    = `equal`
	Contains = `contains`
	Pattern  = `pattern`
)

func IsExpected(expectedType, value string, resp *fasthttp.Response) bool {
	body := string(resp.Body())

	switch expectedType {
	case Contains:
		return strings.Contains(body, value)
	case Pattern:
		ok, _ := regexp.MatchString(value, body)
		return ok
	default:
		return body == value
	}
}
