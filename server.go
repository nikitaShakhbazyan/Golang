package main

import (
	"fmt"
	"net/http"
)

func greeting(name string) string {
	message := fmt.Sprintf("Hi, %v Welcome!", name)
	return message
}

func asking(question string) string {
	message := fmt.Sprintf(question)
	return message
}

func main() {
	myMessage := "Nikita"
	myAsking := "How are you?"
	log := greeting(myMessage)
	askingLog := asking(myAsking)
	fmt.Println(log)
	fmt.Println(askingLog)

	http.HandleFunc("/", server)
	http.HandleFunc("/notMain", notMain)

	fmt.Println("Server is running on host 8081")
	http.ListenAndServe(":8081", nil)
}

func server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")

}
func notMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is not the main page n*ga")
}
