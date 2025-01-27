package main

import (
	"testing"
	"sync"
)

func TestFNV1a(t *testing.T) {
	tests := []struct {
		key    string
		want   uint64
	}{
		{"STUNG", 0x4786a308b6c09372},
		{"", 14695981039346656037},
		{"a", 0xaf63dc4c8601ec8c},
		{"hello", 0xa430d84680aabd0b},
	}

	for _, tt := range tests {
		if got := fnv1a(tt.key); got != tt.want {
			t.Errorf("fnv1a(%q) = %x, want %x", tt.key, got, tt.want)
		}
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		num uint64
		want bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{8796817, true},
		{828372629, false},
		{7638825811, true},
	}

	for _, tt := range tests {
		if got := isPrime(tt.num); got != tt.want {
			t.Errorf("isPrime(%d) = %v, want %v", tt.num, got, tt.want)
		}
	}
}

func TestGenNextPrime(t *testing.T) {
	tests := []struct {
		n     uint64
		want  uint64
	}{
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 5},
		{6, 7},
		{7, 7},
		{8, 11},
		{9, 11},
		{10, 11},
		{11, 11},
		{12, 13},
	}

	for _, tt := range tests {
		if got := genNextPrime(tt.n); got != tt.want {
			t.Errorf("genNextPrime(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestMadHashMultipleTimes(t *testing.T) {
       text := "test"
       expected := madHash(text)
       for i := 0; i < 100; i++ {
              actual := madHash(text)
              if actual != expected {
                     t.Errorf("expected madHash to return %d, got %d", expected, actual)
              }
       }
}


func TestMadHashConcurrency(t *testing.T) {
	text := "test"
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			madHash(text)
		}()
	}

	wg.Wait()
}


func BenchmarkMadHash(b *testing.B) {
	text := "test"
	for i := 0; i < b.N; i++ {
		madHash(text)
	}
}
//*/
