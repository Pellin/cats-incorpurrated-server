package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	data, err := ioutil.ReadFile("data/catdata.json")
	if err != nil {
		fmt.Print(err)
	}

	http.HandleFunc("/cats", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Write(data)
	})

	http.ListenAndServe(":8080", nil)
}
