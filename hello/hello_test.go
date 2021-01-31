// @Author: 2014BDuck
// @Date: 2021/1/31

package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
