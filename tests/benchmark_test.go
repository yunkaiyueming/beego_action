package test

import (
	"beego_action/helpers"
	"testing"
)

func Benchmark_checkMail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(i, b.N, helpers.CheckEmail("819664855@qq.com"))
	}
}
