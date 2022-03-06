package main

import (
	"fmt"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	fmt.Println("Running server v 1.5 Reader")

	http.HandleFunc("/read", read)

	http.ListenAndServe(":8090", nil)
}

func read(w http.ResponseWriter, req *http.Request) {

	dat, err := os.ReadFile("/app/text.txt")
	check(err)
	w.Write(dat)
}
