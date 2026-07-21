package auth

import (
	"strings"
	"testing"
)

// TestEncrypt ensures the encrypted password is a bcrypt hash
// and never equals the plain text.
func TestEncrypt(t *testing.T) {
	hashed, err := Encrypt("my-secret-password")
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}
	if hashed == "my-secret-password" {
		t.Error("Encrypt returned the plain text password")
	}
	if !strings.HasPrefix(hashed, "$2a$") && !strings.HasPrefix(hashed, "$2b$") {
		t.Errorf("Encrypt returned a non-bcrypt hash: %q", hashed)
	}
}

// TestCompare ensures a correct password matches its hash
// and a wrong password is rejected.
func TestCompare(t *testing.T) {
	hashed, err := Encrypt("my-secret-password")
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	if err := Compare(hashed, "my-secret-password"); err != nil {
		t.Errorf("Compare should succeed with the correct password: %v", err)
	}
	if err := Compare(hashed, "wrong-password"); err == nil {
		t.Error("Compare should fail with a wrong password")
	}
	if err := Compare("not-a-hash", "my-secret-password"); err == nil {
		t.Error("Compare should fail with a malformed hash")
	}
}
