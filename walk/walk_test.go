// @Author: 2014BDuck
// @Date: 2021/3/6

package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string field",
			struct {
				Name string
				City string
			}{
				"Chris",
				"London",
			},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non-string field",
			struct {
				Name string
				Age  int
			}{
				"Chris",
				25,
			},
			[]string{"Chris"},
		},
		{
			"Struct with nested field",
			Person{
				Name: "Chris",
				Profile: Profile{
					10,
					"London",
				},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slice to things",
			[]Person{
				{"Chris", Profile{33, "London"}},
				{"Rax", Profile{35, "Guangzhou"}},
			},
			[]string{"Chris", "London", "Rax", "Guangzhou"},
		},
		{
			"array to things",
			[2]Person{
				{"Chris", Profile{33, "London"}},
				{"Rax", Profile{35, "Guangzhou"}},
			},
			[]string{"Chris", "London", "Rax", "Guangzhou"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	assertContains := func(t *testing.T, haystack []string, needle string) {
		contains := false
		for _, x := range haystack {
			if x == needle {
				contains = true
			}
		}
		if !contains {
			t.Errorf("expect %+v to contain '%s' but it didn't", haystack, needle)
		}
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}
