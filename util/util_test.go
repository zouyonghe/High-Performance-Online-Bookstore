package util

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestGenShortId ensures generated short ids are non-empty and unique.
func TestGenShortId(t *testing.T) {
	seen := make(map[string]struct{})
	for i := 0; i < 100; i++ {
		id, err := GenShortId()
		if err != nil {
			t.Fatalf("GenShortId failed: %v", err)
		}
		if id == "" {
			t.Fatal("GenShortId returned an empty id")
		}
		if _, ok := seen[id]; ok {
			t.Fatalf("GenShortId returned a duplicated id: %q", id)
		}
		seen[id] = struct{}{}
	}
}

// TestGetReqID ensures the request id is extracted from the gin context.
func TestGetReqID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	if got := GetReqID(c); got != "" {
		t.Errorf("GetReqID without context value = %q, want empty", got)
	}

	c.Set("X-Request-Id", "req-123")
	if got := GetReqID(c); got != "req-123" {
		t.Errorf("GetReqID = %q, want %q", got, "req-123")
	}
}
