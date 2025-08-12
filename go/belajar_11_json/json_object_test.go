package belajar11json

import (
	"encoding/json"
	"testing"
)

type Customer struct {
	FirstName string
	MiddleName string 
	LastName  string
	Age int
	Married bool
	Hobbies []string
}

func TestJson(t *testing.T){
	customer := Customer{
		FirstName: "John",
		MiddleName: "Doe",
		LastName:  "Smith",
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		t.Fatalf("Failed to marshal customer: %v", err)
	}

	expected := `{"FirstName":"John","MiddleName":"Doe","LastName":"Smith"}`
	if string(bytes) != expected {
		t.Errorf("Expected %s but got %s", expected, string(bytes))
	}
}