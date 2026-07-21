package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"High-Performance-Online-Bookstore/pkg/berror"

	"github.com/gin-gonic/gin"
)

// TestSendResponse ensures a nil error produces code 0 with data.
func TestSendResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	SendResponse(c, nil, map[string]string{"hello": "world"})

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
	}
	var rsp Response
	if err := json.Unmarshal(w.Body.Bytes(), &rsp); err != nil {
		t.Fatalf("response is not valid JSON: %v", err)
	}
	if rsp.Code != berror.OK.Code {
		t.Errorf("code = %d, want %d", rsp.Code, berror.OK.Code)
	}
}

// TestSendError ensures a business error is decoded into the response.
func TestSendError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	SendError(c, berror.ErrUserNotFound)

	var rsp Response
	if err := json.Unmarshal(w.Body.Bytes(), &rsp); err != nil {
		t.Fatalf("response is not valid JSON: %v", err)
	}
	if rsp.Code != berror.ErrUserNotFound.Code {
		t.Errorf("code = %d, want %d", rsp.Code, berror.ErrUserNotFound.Code)
	}
}

// TestSendDenyResponse ensures a denied request returns 403.
func TestSendDenyResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	SendDenyResponse(c)

	if w.Code != http.StatusForbidden {
		t.Errorf("status = %d, want %d", w.Code, http.StatusForbidden)
	}
}

// TestNoRoute ensures unknown routes return 404.
func TestNoRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	NoRoute(c)

	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want %d", w.Code, http.StatusNotFound)
	}
}
