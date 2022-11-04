package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type MathOperationsReq struct {
	OperationType string `json:"operation_type"`
	X             int    `json:"x"`
	Y             int    `json:"y"`
}

type MathOperationResult struct {
	SlackUsername string `json:"slackUsername"`
	Result        int    `json:"result"`
	OperationType string `json:"operation_type"`
}

func mathHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	req := &MathOperationsReq{}

	if err := dec.Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := &MathOperationResult{
		SlackUsername: "Bolu_adx",
		OperationType: req.OperationType,
	}
	switch req.OperationType {
	case "addition":
		resp.Result = req.X + req.Y
	case "subtraction":
		resp.Result = req.X - req.Y
	case "multiplication":
		resp.Result = req.X * req.Y
	}
	w.Header().Set("content-Type", "application/json")

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		log.Printf("cannot encode %v - %s", resp, err)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/mathOperations", mathHandler).Methods("POST")

	fmt.Printf("starting server at port :8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
