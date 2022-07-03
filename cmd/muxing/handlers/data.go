package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type DataHandler struct{}

func (h DataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	param, ok := vars["param"]
	w.WriteHeader(http.StatusOK)

	if ok {
		fmt.Fprintf(w, "Hello, %s!", param)
	}
}
