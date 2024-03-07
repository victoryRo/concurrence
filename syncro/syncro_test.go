package syncro

import "testing"

func BenchmarkTimerConcurrent(b *testing.B) {
	times := []int{1, 2, 5, 7}

	for i := 0; i < b.N; i++ {
		timerConcurrent(times)
	}
}

func BenchmarkTimer(b *testing.B) {
	times := []int{2, 4, 5}

	for i := 0; i < b.N; i++ {
		timer(times)
	}
}
