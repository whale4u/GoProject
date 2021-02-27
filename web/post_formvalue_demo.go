package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var (
				name string = r.PostFormValue("name")
				id   string = r.PostFormValue("id")
			)
			fmt.Printf("name-id is  : %s %s\n", name, id)
		}
	})

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
