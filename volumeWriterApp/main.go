package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	fmt.Println("Running server v 1.5 Writer")

	http.HandleFunc("/write", write)

	http.ListenAndServe(":8091", nil)
}

func write(w http.ResponseWriter, req *http.Request) {

	path := "/app/text.txt"

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		if err.Error() == fmt.Sprintf("open %s: no such file or directory", path) {
			f, err = os.Create(path)
			check(err)
		}
		check(err)
	}

	f.WriteString(fmt.Sprintf("%b", rand.Intn(100)) + "\n")
}
