package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GolangUnited/helloweb/cmd/muxing/handlers"
	"github.com/gorilla/mux"
)

func TestDataEndpoint(t *testing.T) {
	cases := []string{"test", "hello", "world"}

	for _, test := range cases {
		t.Run("case "+test, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/name/%s", test), nil)
			w := httptest.NewRecorder()
			h := handlers.DataHandler{}

			router := mux.NewRouter()
			router.Handle("/name/{param}", h)
			router.ServeHTTP(w, r)

			result := w.Result()
			defer result.Body.Close()
			data, err := ioutil.ReadAll(result.Body)

			if err != nil {
				t.Error(err)
			}

			expected := fmt.Sprintf("Hello, %s!", test)

			if string(data) != expected {
				t.Errorf("Expected `%s`, got `%s`", expected, string(data))
			}

			if w.Code != http.StatusOK {
				t.Errorf("Expected `%d`, got `%d`", http.StatusOK, w.Code)
			}
		})
	}
}
