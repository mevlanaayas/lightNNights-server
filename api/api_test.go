package api

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
var api *API
var testConfig *config.Config
*/

func TestCanSayHello(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, hello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello, World!", rec.Body.String())
	}
}

/*
func TestInfoEndpoint(t *testing.T) {
	code, body := request(t, "GET", "/info", nil)
	if assert.Equal(t, http.StatusOK, code) {
		raw := make(map[string]string)
		extractPayload(t, body, &raw)
		assert.NotEmpty(t, raw["version"])
		assert.NotEmpty(t, raw["description"])
		assert.NotEmpty(t, raw["name"])
	}
}

func extractPayload(t *testing.T, body string, out interface{}) {
	err := json.Unmarshal([]byte(body), out)
	assert.NoError(t, err)
}

func request(t *testing.T, method, path string, body interface{}) (int, string) {
	req := httptest.NewRequest(method, path, nil)

	if body != nil {
		bs, err := json.Marshal(body)
		if err != nil {
			assert.FailNow(t, "failed to serialize request body: "+err.Error())
		}

		req = httptest.NewRequest(method, path, bytes.NewBuffer(bs))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	}

	rsp := httptest.NewRecorder()
	api.echo.ServeHTTP(req, rsp)
	return rsp.Code, rsp.Body.String()
}
*/
