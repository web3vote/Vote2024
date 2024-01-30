package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RequestBody struct {
    MRZ string `json:"mrz"`
   }
   
   func ProcessMRZ(w http.ResponseWriter, r *http.Request) {
    // Read the HTTP request body
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
     http.Error(w, "Error reading request body", http.StatusBadRequest)
     return
    }

    // Parse the JSON body into the defined struct
    var requestBody RequestBody
    err = json.Unmarshal(body, &requestBody)
    if err != nil {
     http.Error(w, "Error parsing JSON body", http.StatusBadRequest)
     return
    }

    // Validate MRZ using TTP API (replace with your actual TTP API logic)
    isValid := ValidateMRZ(requestBody.MRZ)

    // Get Keccak hash from MRZ (replace with your actual logic)
    keccakHash := GetKeccakHash(requestBody.MRZ)

    // Submit hash to Solidity contract (replace with your actual contract interaction)
    submitResult := SubmitToContract(keccakHash)

    // Prepare the response JSON
    response := struct {
     Valid       bool   `json:"valid"`
     KeccakHash  string `json:"keccakHash"`
     SubmitResult string `json:"submitResult"`
    }{
     Valid:       isValid,
     KeccakHash:  keccakHash,
     SubmitResult: submitResult,
    }

    // Convert the response to JSON
    responseJSON, err := json.Marshal(response)
    if err != nil {
     http.Error(w, "Error creating response JSON", http.StatusInternalServerError)
    }

    // Set the response headers
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    // Send the response
    fmt.Fprint(w, string(responseJSON))
   }
   

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/process-mrz", ProcessMRZ).Methods("POST")

    log.Fatal(http.ListenAndServe(":8000", router))
   }
   


