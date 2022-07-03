package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/GolangUnited/helloweb/cmd/muxing/handlers"
	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	dataHandler := &handlers.DataHandler{}

	router.Handle("/name/{param}", dataHandler).Methods(http.MethodGet)

	router.HandleFunc("/bad", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}).Methods(http.MethodGet)

	router.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "I got message:\n%s", data)
	}).Methods(http.MethodPost)

	router.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		a := r.Header.Get("a")
		b := r.Header.Get("b")

		aNum, err := strconv.Atoi(a)

		if err != nil {
			return
		}

		bNum, err := strconv.Atoi(b)

		if err != nil {
			return
		}

		w.Header().Add("a+b", strconv.Itoa(aNum+bNum))
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodPost)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
