package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerGetRR(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Error(err)
	}
	handleGet(rr, req)
	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, rr.Result().StatusCode)
	}

	expected := "hello world"
	b, err := io.ReadAll(rr.Result().Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("expected %s, got %s", expected, string(b))
	}
}

func TestHandlerGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handleGet))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, resp.StatusCode)
	}

	expected := "hello world"
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("expected %s, got %s", expected, string(b))
	}
}
