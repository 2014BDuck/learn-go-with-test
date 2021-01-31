// @Author: 2014BDuck
// @Date: 2021/1/31

package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if want != got {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})

	t.Run("sum slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		got := Sum(numbers)
		want := 21
		if want != got {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4, 5})
	want := []int{3, 12}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumTail(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumTail([]int{1, 2, 3, 4, 5}, []int{1, 100})
		want := []int{14, 100}
		checkSums(t, got, want)
	})

	t.Run("make the sums of empty slice", func(t *testing.T) {
		got := SumTail([]int{1, 2, 3, 4, 5}, []int{})
		want := []int{14, 0}
		checkSums(t, got, want)
	})

}
