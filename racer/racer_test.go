// @Author: 2014BDuck
// @Date: 2021/3/6

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("test racer", func(t *testing.T) () {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL, 10*time.Second)

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("test timeout", func(t *testing.T) () {
		var timeoutDuration = 1 * time.Second
		serverA := makeDelayedServer(timeoutDuration + 1*time.Second)
		serverB := makeDelayedServer(timeoutDuration + 2*time.Second)
		defer serverA.Close()
		defer serverB.Close()

		slowURL := serverA.URL
		fastURL := serverB.URL

		_, err := Racer(slowURL, fastURL, timeoutDuration)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
