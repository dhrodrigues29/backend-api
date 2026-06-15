package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJSONSetsContentType(t *testing.T) {
	rec := httptest.NewRecorder()

	payload := map[string]string{
		"status": "ok",
	}

	writeJSON(rec, http.StatusOK, payload)

	contentType := rec.Header().Get("Content-Type")

	if contentType != "application/json" {
		t.Fatalf(
			"expected Content-Type application/json, got %s",
			contentType,
		)
	}
}

func TestWriteJSONSetsStatusCode(t *testing.T) {
	rec := httptest.NewRecorder()

	writeJSON(
		rec,
		http.StatusCreated,
		map[string]string{"status": "created"},
	)

	if rec.Code != http.StatusCreated {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusCreated,
			rec.Code,
		)
	}
}

func TestWriteJSONWritesBody(t *testing.T) {
	rec := httptest.NewRecorder()

	expected := map[string]string{
		"status": "ok",
	}

	writeJSON(rec, http.StatusOK, expected)

	var actual map[string]string

	if err := json.Unmarshal(rec.Body.Bytes(), &actual); err != nil {
		t.Fatalf("failed to parse response body: %v", err)
	}

	if actual["status"] != expected["status"] {
		t.Fatalf(
			"expected status %s, got %s",
			expected["status"],
			actual["status"],
		)
	}
}