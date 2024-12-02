package main

import (
	"fmt"
	"net/http"
)

func handlerHttp (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"HTTP")
}

func main () {

	http.HandleFunc("/main",handlerHttp)

	fmt.Println("server is started on port 8080")

	if err := http.ListenAndServe(":8080",nil); err != nil {
		fmt.Println("Error starting server",err)
	}

}