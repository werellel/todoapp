package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMux(t *testing.T) {
	// httptest.NewRecorder 함수를 사용하면 ResponseWriter 인터페이스를
	// 충족하는 Response Recorder 타입의 값을 얻을 수 있다.
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	sut := NewMux()
	sut.ServeHTTP(w, r)
	resp := w.Result()
	t.Cleanup(func() { _ = resp.Body.Close() })
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	want := `{"status": "ok"}`
	if string(got) != want {
		t.Errorf("unexpected response body: %s", got)
	}

}
