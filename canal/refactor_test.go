package canal

import "testing"

func BenchmarkTimer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		timer(items(b))
	}
}

func BenchmarkTimerConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		timerConcurrent(items(b))
	}
}

func BenchmarkTimerChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		timerChannelCSP(items(b))
	}
}

func items(b *testing.B) []int {
	b.Helper()
	return []int{1, 2, 5}
}
