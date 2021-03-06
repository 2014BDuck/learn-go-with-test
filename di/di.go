// @Author: 2014BDuck
// @Date: 2021/3/6

package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	_, err := fmt.Fprintf(writer, "Hello, %s", name)
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
	if err != nil{
		fmt.Printf("%v", err.Error())
	}
}
