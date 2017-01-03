package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todos := Todos {
		Todo {Name: "Hoge"},
		Todo {Name: "Fuga"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Todo show: %s", p.ByName("todoId"))
}
