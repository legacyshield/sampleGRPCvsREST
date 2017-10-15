package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type ExampleProfileRequest struct {
	ExampleName string `json:"name"`
	ExampleAge 	int64 `json:"age"`
	ExampleCity string `json:"city"`
	ExampleCountry string `json:"country"`
}

type ExampleProfileCreationResponse struct {
	Message string `json:"message"`
}

func createProfile(w http.ResponseWriter, r *http.Request) {
	var profile ExampleProfileRequest
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		fmt.Println("error in decoding ", err)
	}
	fmt.Println(profile)

	response := ExampleProfileCreationResponse{
		Message: "Profile Created with name" + profile.ExampleName,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(response)
}
func main() {
	http.HandleFunc("/profile", createProfile)
	http.ListenAndServe(":3000", nil)
}
