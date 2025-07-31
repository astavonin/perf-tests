package dailytemps

import (
	"flag"
	"reflect"
	"testing"
	"time"
)

func TestDailyTemperatures(t *testing.T) {
	input := []int{73, 74, 75, 71, 69, 72, 76, 73}
	expected := []int{1, 1, 4, 2, 1, 1, 0, 0}

	got := DailyTemperatures(input)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("DailyTemperatures failed: expected %v, got %v", expected, got)
	}

	gotOpt := DailyTemperaturesOpt(input)
	if !reflect.DeepEqual(gotOpt, expected) {
		t.Errorf("DailyTemperaturesOpt failed: expected %v, got %v", expected, gotOpt)
	}
}

var benchSize = flag.Int("bench.size", 10000, "Input size for benchmark")

func generateInput(n int) []int {
	input := make([]int, n)
	for i := range input {
		input[i] = 30 + i%40
	}
	return input
}

func BenchmarkDailyTemperatures(b *testing.B) {
	input := generateInput(*benchSize)

	b.Run("Baseline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = DailyTemperatures(input)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = DailyTemperaturesOpt(input)
		}
	})
}

func BenchmarkTracePerRun(b *testing.B) {
	input := generateInput(*benchSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		start := time.Now()
		_ = DailyTemperaturesOpt(input)
		b.Logf("Run %d: %v for %d elements", i+1, time.Since(start), *benchSize)
	}
}
