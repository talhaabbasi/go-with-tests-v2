package sum

import (
	"fmt"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1,2,3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d but want %d, %v", got, want, numbers)
		}
	})
}

func ExampleSum() {
	sum := Sum([]int{1,2,3})
	fmt.Println(sum)
	// Output: 6
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1,2,3})
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1,2,3}, []int{0, 9})
		want := []int{5, 9}

		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

func ExampleSumAllTails() {
	sum := SumAllTails([]int{1,2,3})
	fmt.Println(sum)
	// Output: [5]
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1,2,3})
	}
}

func TestSumAllHeads(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		got := SumAllHeads([]int{1,2,3}, []int{0, 9})
		want := []int{1, 0}

		if !slices.Equal(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})
}

func ExampleSumAllHeads() {
	sum := SumAllHeads([]int{1,2,3})
	fmt.Println(sum)
	// Output: [1]
}

func BenchmarkSumAllHeads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1,2,3})
	}
}
