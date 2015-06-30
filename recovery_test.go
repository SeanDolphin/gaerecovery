package gaerecovery

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codegangsta/negroni"
)

func TestRecovery(t *testing.T) {

	recorder := httptest.NewRecorder()

	rec := Recovery()

	n := negroni.New()
	n.Use(rec)
	n.UseHandler(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		panic("here is a panic!")
	}))
	n.ServeHTTP(recorder, (*http.Request)(nil))

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Expected %v - Got %v", http.StatusInternalServerError, recorder.Code)
	}
}
