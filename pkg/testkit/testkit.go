package testkit

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func DoEchoRequest(handler echo.HandlerFunc, req *http.Request) *httptest.ResponseRecorder {
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

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
