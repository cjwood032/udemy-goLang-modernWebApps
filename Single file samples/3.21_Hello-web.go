package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hi Mom!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v bytes written", n)
	})
	_ = http.ListenAndServe(":3000", nil)
}
