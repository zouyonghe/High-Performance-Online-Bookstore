package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const testSecret = "test-secret-key-for-jwt-signing"

// TestSignAndParse signs a fresh token and parses it back,
// verifying that all claims survive the round trip.
func TestSignAndParse(t *testing.T) {
	c := Context{
		ID:       123,
		Username: "test",
		Role:     "general",
	}
	tokenString, err := Sign(c, testSecret)
	if err != nil {
		t.Fatalf("Sign failed: %v", err)
	}

	ctx, err := Parse(tokenString, testSecret)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	if ctx.ID != c.ID || ctx.Username != c.Username || ctx.Role != c.Role {
		t.Errorf("claims mismatch: got %+v, want %+v", ctx, c)
	}
}

// TestParseWrongSecret ensures a token signed with another
// secret is rejected.
func TestParseWrongSecret(t *testing.T) {
	tokenString, err := Sign(Context{ID: 1, Username: "u", Role: "general"}, testSecret)
	if err != nil {
		t.Fatalf("Sign failed: %v", err)
	}
	if _, err = Parse(tokenString, "another-secret"); err == nil {
		t.Error("Parse should fail with a wrong secret")
	}
}

// TestParseExpired ensures an expired token is rejected.
func TestParseExpired(t *testing.T) {
	expired := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       123,
		"username": "test",
		"role":     "general",
		"nbf":      time.Now().Add(-2 * time.Hour).Unix(),
		"iat":      time.Now().Add(-2 * time.Hour).Unix(),
		"exp":      time.Now().Add(-time.Hour).Unix(),
	})
	tokenString, err := expired.SignedString([]byte(testSecret))
	if err != nil {
		t.Fatalf("SignedString failed: %v", err)
	}
	if _, err = Parse(tokenString, testSecret); err == nil {
		t.Error("Parse should fail with an expired token")
	}
}

// TestParseMalformedClaims ensures a token with malformed
// claims does not panic and returns an error.
func TestParseMalformedClaims(t *testing.T) {
	malformed := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       "not-a-number",
		"username": "test",
		"role":     "general",
		"exp":      time.Now().Add(time.Hour).Unix(),
	})
	tokenString, err := malformed.SignedString([]byte(testSecret))
	if err != nil {
		t.Fatalf("SignedString failed: %v", err)
	}
	if _, err = Parse(tokenString, testSecret); err == nil {
		t.Error("Parse should fail with malformed claims")
	}
}

// TestParseAlgConfusion ensures tokens signed with an
// unexpected algorithm (e.g. none) are rejected.
func TestParseAlgConfusion(t *testing.T) {
	noneToken := jwt.NewWithClaims(jwt.SigningMethodNone, jwt.MapClaims{
		"id":       123,
		"username": "test",
		"role":     "admin",
		"exp":      time.Now().Add(time.Hour).Unix(),
	})
	tokenString, err := noneToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	if err != nil {
		t.Fatalf("SignedString failed: %v", err)
	}
	if _, err = Parse(tokenString, testSecret); err == nil {
		t.Error("Parse should reject tokens using the none algorithm")
	}
}
