package test

import (
	"beego_action/helpers"
	"testing"
)

func TestCheckEmail(t *testing.T) {
	if helpers.CheckEmail("845688@qq.com") {
		t.Log("ok")
	} else {
		t.Error("failed")
	}
}
