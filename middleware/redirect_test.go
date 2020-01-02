package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jimyx17/echo/v4"
	"github.com/stretchr/testify/assert"
)

type middlewareGenerator func() echo.MiddlewareFunc

func TestRedirectHTTPSRedirect(t *testing.T) {
	res := redirectTest(HTTPSRedirect, "jimyx17.com", nil)

	assert.Equal(t, http.StatusMovedPermanently, res.Code)
	assert.Equal(t, "https://jimyx17.com/", res.Header().Get(echo.HeaderLocation))
}

func TestHTTPSRedirectBehindTLSTerminationProxy(t *testing.T) {
	header := http.Header{}
	header.Set(echo.HeaderXForwardedProto, "https")
	res := redirectTest(HTTPSRedirect, "jimyx17.com", header)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestRedirectHTTPSWWWRedirect(t *testing.T) {
	res := redirectTest(HTTPSWWWRedirect, "jimyx17.com", nil)

	assert.Equal(t, http.StatusMovedPermanently, res.Code)
	assert.Equal(t, "https://www.jimyx17.com/", res.Header().Get(echo.HeaderLocation))
}

func TestRedirectHTTPSWWWRedirectBehindTLSTerminationProxy(t *testing.T) {
	header := http.Header{}
	header.Set(echo.HeaderXForwardedProto, "https")
	res := redirectTest(HTTPSWWWRedirect, "jimyx17.com", header)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestRedirectHTTPSNonWWWRedirect(t *testing.T) {
	res := redirectTest(HTTPSNonWWWRedirect, "www.jimyx17.com", nil)

	assert.Equal(t, http.StatusMovedPermanently, res.Code)
	assert.Equal(t, "https://jimyx17.com/", res.Header().Get(echo.HeaderLocation))
}

func TestRedirectHTTPSNonWWWRedirectBehindTLSTerminationProxy(t *testing.T) {
	header := http.Header{}
	header.Set(echo.HeaderXForwardedProto, "https")
	res := redirectTest(HTTPSNonWWWRedirect, "www.jimyx17.com", header)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestRedirectWWWRedirect(t *testing.T) {
	res := redirectTest(WWWRedirect, "jimyx17.com", nil)

	assert.Equal(t, http.StatusMovedPermanently, res.Code)
	assert.Equal(t, "http://www.jimyx17.com/", res.Header().Get(echo.HeaderLocation))
}

func TestRedirectNonWWWRedirect(t *testing.T) {
	res := redirectTest(NonWWWRedirect, "www.jimyx17.com", nil)

	assert.Equal(t, http.StatusMovedPermanently, res.Code)
	assert.Equal(t, "http://jimyx17.com/", res.Header().Get(echo.HeaderLocation))
}

func redirectTest(fn middlewareGenerator, host string, header http.Header) *httptest.ResponseRecorder {
	e := echo.New()
	next := func(c echo.Context) (err error) {
		return c.NoContent(http.StatusOK)
	}
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Host = host
	if header != nil {
		req.Header = header
	}
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	fn()(next)(c)

	return res
}
