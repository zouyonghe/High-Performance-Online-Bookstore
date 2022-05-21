package token

import (
	"fmt"
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
	_, err := Sign(c, secret)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("TestSign passed")
	}
}

func TestParse(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTMzNjkxODIsImlhdCI6MTY1MzEwOTk4MiwiaWQiOjEyMywibmJmIjoxNjUzMTA5OTgyLCJyb2xlIjoiZ2VuZXJhbCIsInVzZXJuYW1lIjoidGVzdCJ9.9a6HR2sZHq9R9wRicGRUmJaKv_yhtdFNJ64fe739Etg"
	c, err := Parse(token, viper.GetString("jwt_secret"))
	if err != nil {
		t.Error(err)
	}
	if c.ID == 123 && c.Username == "test" && c.Role == "general" {
		fmt.Println("TestParse passed")
	} else {
		t.Error("TestParse failed")
	}
}
