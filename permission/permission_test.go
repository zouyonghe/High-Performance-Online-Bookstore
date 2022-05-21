package permission

import (
	"fmt"
	"testing"
)

type testCase struct {
	sub string
	obj string
	act string
}

func TestCheckPermission(t *testing.T) {
	E = initPermission("../conf/model.conf", "../conf/policy.csv", false)
	test1 := testCase{
		sub: "admin",
		obj: "/v1/user/admin",
		act: "GET",
	}
	ok1 := CheckPermission(test1.sub, test1.obj, test1.act)
	test2 := testCase{
		sub: "general",
		obj: "/v1/user/admin",
		act: "POST",
	}
	ok2 := CheckPermission(test2.sub, test2.obj, test2.act)
	if ok1 && !ok2 {
		fmt.Println("Test CheckPermission passed")
	} else {
		t.Error("Test CheckPermission failed")
	}
}
