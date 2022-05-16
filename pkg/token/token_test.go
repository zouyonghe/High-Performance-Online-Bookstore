package token

import (
	"github.com/spf13/viper"
	"testing"
)

func TestSign(t *testing.T) {
	c := Context{
		ID:       123,
		Username: "test",
		Role:     "general",
	}
	secret := viper.GetString("jwt_secret")
	token, err := Sign(c, secret)
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

func TestParse(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTI5NjMzNzUsImlhdCI6MTY1MjcwNDE3NSwiaWQiOjEyMywibmJmIjoxNjUyNzA0MTc1LCJyb2xlIjoiZ2VuZXJhbCIsInVzZXJuYW1lIjoidGVzdCJ9.jD1JC2ulNQLS0d1b4R_pV27N-9jn0PiJ153OdwUUiok"
	c, err := Parse(token, viper.GetString("jwt_secret"))
	if err != nil {
		t.Error(err)
	}
	t.Log(c)
}
