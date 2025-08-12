package belajar11json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	customerJson := `{"FirstName":"John","MiddleName":"Doe","LastName":"Smith"}`
	jsonBytes := []byte(customerJson)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)

	
}