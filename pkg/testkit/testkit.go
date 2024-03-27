package testkit

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

type doEchoRequestOptionFunc func(*doEchoRequestOption)

type doEchoRequestOption struct {
	params map[string]string
}

func WithParams(params map[string]string) doEchoRequestOptionFunc {
	return func(o *doEchoRequestOption) {
		o.params = params
	}
}

func bindOptions(opts []doEchoRequestOptionFunc) *doEchoRequestOption {
	o := &doEchoRequestOption{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func DoEchoRequest(handler echo.HandlerFunc, req *http.Request, opts ...doEchoRequestOptionFunc) *httptest.ResponseRecorder {
	o := bindOptions(opts)
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for k, v := range o.params {
		c.SetParamNames(k)
		c.SetParamValues(v)
	}

	handler(c)

	return rec
}

func JSONStringify(t *testing.T, v any) string {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("unable to marshal json: %v", err)
	}
	return string(b)
}
