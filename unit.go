package fasthttpunit

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type Unit struct {
	casePath string
	conf     *Conf
	t        *testing.T
	server   *fasthttp.Server
}

func NewUnit(conf *Conf, t *testing.T, s *fasthttp.Server) *Unit {
	u := &Unit{
		t:      t,
		server: s,
		conf:   conf,
	}

	return u
}

func NewUnitWithHandler(conf *Conf, t *testing.T, handler fasthttp.RequestHandler) *Unit {
	s := &fasthttp.Server{
		Handler: handler,
	}
	return NewUnit(conf, t, s)
}

func NewUnitWithRouter(conf *Conf, t *testing.T, r *fasthttprouter.Router) *Unit {
	return NewUnitWithHandler(conf, t, r.Handler)
}

func (u *Unit) Test(mockList ...func()) {
	u.Green("--UNIT TEST WILL START AT %d SECONDS LATER--", u.conf.Delay)
	delay := time.Duration(u.conf.Delay) * time.Second
	time.Sleep(delay)

	if len(mockList) > 0 {
		for _, mock := range mockList {
			funcName := runtime.FuncForPC(reflect.ValueOf(mock).Pointer()).Name()
			u.Green("START RUN mock:%s", funcName)
			mock()
		}

		u.Green("---RUN mock done---")
	}

	u.Green("----START UNIT TEST----")

	for _, api := range u.conf.ApiList {
		u.Green("--------START TEST API:【%s】--------", api.Desc)

		for _, cs := range api.CaseList {
			u.testCase(api, cs)
		}

		fmt.Println()
	}
}

func (u *Unit) testCase(api *Api, cs Case) {
	var (
		req  = CaseToRequest(api, cs)
		resp = fasthttp.AcquireResponse()
	)

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	rw := NewReadWriter(u.server)

	u.Green("@@@@@ START RUN CASE:【%s】 %s", cs.Desc, strings.TrimSuffix(req.String(), "\n"))
	err := rw.Request(req, resp)
	if err != nil {
		u.Fatal("REQUEST ERROR:%s", err)
	}

	if !IsExpected(cs.ExpectedType, cs.Expected, resp) {
		u.Red("###########EXECUTE API【%s】 CASE：【%s】FAILED###########", api.Desc, cs.Desc)
		u.Yellow("RESPONSE STATUS:%d BODY:%s\n",
			resp.StatusCode(),
			resp.Body(),
		)
	} else {
		u.Green("###########PASS###########\n")
	}
}

func (u *Unit) Fatal(format string, args ...interface{}) {
	u.t.Fatal(Red(format, args...))
}

func (u *Unit) Red(format string, args ...interface{}) {
	u.t.Logf(Red(format, args...))
}

func (u *Unit) Green(format string, args ...interface{}) {
	u.t.Logf(Green(format, args...))
}

func (u *Unit) Yellow(format string, args ...interface{}) {
	u.t.Logf(Yellow(format, args...))
}
