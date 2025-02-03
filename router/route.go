package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Route() {
	r := mux.NewRouter()

	r.HandleFunc("/api/classify-number/{number}", func(w http.ResponseWriter, r *http.Request) {
		variables := mux.Vars(r)
		number := variables["number"]

		fmt.Fprintf(w, "This is your number %s", number)
	})
	log.Println("Server on port :3000----------------")
	log.Fatal(http.ListenAndServe(":3000", r))
}
