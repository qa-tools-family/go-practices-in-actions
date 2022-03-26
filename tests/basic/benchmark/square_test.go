package math

import "testing"

func BenchmarkSquare(b *testing.B) {
	// 可以执行一些准备工作，不计入耗时
	b.ResetTimer()
	for i := 0; i < b.N; i++ {  // b.N 为运行时动态计算得到的运行次数
		Square(3)
	}
	b.StopTimer()
	// 可以执行一些释放资源的函数，不计入函数
}
