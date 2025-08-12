package belajar11json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T){
	customer := Customer{
		FirstName: "John",
		MiddleName: "Doe",
		LastName:  "Smith",
		Age: 30,
		Married: true,
		Hobbies: []string{"Reading", "Traveling", "Cooking"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		t.Fatalf("Failed to marshal customer: %v", err)
	}

	fmt.Println(string(bytes))
}