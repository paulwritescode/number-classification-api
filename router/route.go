package router

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/paulwritescode/numbers-api/analysis"
	"log"
	"net/http"
	"strconv"
)

type ErrorMessage struct {
	Number string `json:"number"`
	Error  bool   `json:"error"`
}

func corsMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") //allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		//handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Route() {
	r := mux.NewRouter()
	r.Use(corsMiddleWare)
	r.HandleFunc("/api/classify-number", classifyNumber).Methods("Get")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func classifyNumber(w http.ResponseWriter, r *http.Request) {
	queryNumber := r.URL.Query()
	queriedNumber := queryNumber.Get("number") //get the number from query params

	// check if the number exsits
	if queriedNumber == "" {
		http.Error(w, "Missing 'number' query parameter", http.StatusBadRequest)
		return
	}

	//convert the queried number to integer
	intQueriedNumber, err := strconv.Atoi(queriedNumber)
	if err != nil {
		error_message := ErrorMessage{
			Number: "alphabet",
			Error:  true,
		}

		ERRORMESSAGE, err := json.MarshalIndent(error_message, "", "  ")
		if err != nil {
			http.Error(w, "Something is wrong with the error handling", http.StatusBadRequest)
		} else {
			http.Error(w, string(ERRORMESSAGE), http.StatusBadRequest)
		}
		return
	}

	JSONRETURN := analysis.ReturnNumber(intQueriedNumber)

	// encode to json
	JsonReturn, err := json.MarshalIndent(JSONRETURN, "", "")
	if err != nil {
		http.Error(w, "Error converting to Json", http.StatusBadRequest)
	}
	fmt.Fprint(w, string(JsonReturn))
}
