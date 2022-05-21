package common

import (
	"High-Performance-Online-Bookstore/model"
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {
	u1 := model.User{
		Username: "test",
		Password: "123456",
		Role:     "general",
	}
	err1 := u1.Validate()
	u2 := model.User{
		Username: "test",
		Password: "123",
		Role:     "general",
	}
	err2 := u2.Validate()
	u3 := model.User{
		Username: "test",
		Password: "123456",
		Role:     "admin",
	}
	err3 := u3.Validate()
	if err1 == nil && err2 != nil && err3 != nil {
		fmt.Println("TestValidate pass")
	} else {
		t.Error("TestValidate fail")
	}
}
