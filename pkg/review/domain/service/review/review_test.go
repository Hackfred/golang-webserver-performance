package review

import (
	"testing"
)

const N = 1000

type payload struct {
	v [N]uint8
}

func doSomeWork(in payload, repeat int) payload {
	if repeat == 0 {
		return in
	}

	return doSomeWork(in, repeat-1)
}

func doSomeWorkRef(in *payload, repeat int) *payload {
	if repeat == 0 {
		return in
	}

	return doSomeWorkRef(in, repeat-1)
}

func BenchmarkWork10(b *testing.B) {
	runBench(10, b)
}

func BenchmarkWork100(b *testing.B) {
	runBench(100, b)
}

func BenchmarkWork1000(b *testing.B) {
	runBench(1000, b)
}

func BenchmarkWorkRef10(b *testing.B) {
	runBenchRef(10, b)
}

func BenchmarkWorkRef100(b *testing.B) {
	runBenchRef(100, b)
}

func BenchmarkWorkRef1000(b *testing.B) {
	runBenchRef(1000, b)
}

func runBench(reps int, b *testing.B) {
	in := payload{v: [N]uint8{}}

	for i := 0; i < b.N; i++ {
		doSomeWork(in, reps)
	}
}

func runBenchRef(reps int, b *testing.B) {
	in := &payload{v: [N]uint8{}}

	for i := 0; i < b.N; i++ {
		doSomeWorkRef(in, reps)
	}
}
