package main

import (
	"net/http/httptest"
	"testing"
)

func TestSink(t *testing.T) {
	for _, path := range []string{"/", "/nsx/event", "/anything/at/all?x=1"} {
		rec := httptest.NewRecorder()
		handler().ServeHTTP(rec, httptest.NewRequest("POST", path, nil))
		if rec.Code != 200 || rec.Body.Len() != 0 {
			t.Fatalf("%s: code %d body %q", path, rec.Code, rec.Body.String())
		}
	}
	rec := httptest.NewRecorder()
	handler().ServeHTTP(rec, httptest.NewRequest("GET", "/healthz", nil))
	if rec.Code != 200 || rec.Body.String() != "ok" {
		t.Fatalf("healthz %d %q", rec.Code, rec.Body.String())
	}
}
