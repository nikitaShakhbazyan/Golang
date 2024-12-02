package main

import (
	"fmt"
	"net/http"
)


func greeting (name string) string {
    message := fmt.Sprintf("Hi, %v Welcome!", name)
	return message
}

func main () {
	myMessage := "Nikita"
	log := greeting(myMessage)
	fmt.Println(log)

	http.HandleFunc("/",server)
	http.HandleFunc("/notMain",notMain)

	fmt.Println("Server is running on host 8081")
	http.ListenAndServe(":8081",nil)
}

func server (w http.ResponseWriter , r *http.Request) {
	fmt.Fprintln(w,"Hello world")

}
func notMain (w http.ResponseWriter , r *http.Request) {
	fmt.Fprintln(w,"This is not the main page n*ga")
}