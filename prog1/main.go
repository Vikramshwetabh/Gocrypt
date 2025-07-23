package main

import (
	"encoding/json" //  package for JSON handling
	"fmt"
	"os"
)

type Person struct {
	Name          string `json:"name"`
	Mobile        string `json:"mobile"`
	BloodGroup    string `json:"blood_group"`
	Email         string `json:"email"`
	AadhaarNumber string `json:"aadhaar_number"`
}

type Data struct {
	Person Person `json:"person"` //struct to hold person data
}

func main() {
	fileData, err := os.ReadFile("input.json")
	if err != nil {
		panic(err)
	}

	var data Data // Create an instance of Data to hold the unmarshalled data
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		panic(err)
	}

	fmt.println("Original Data:")
	fmt.Println("Name:", data.Person.Name)
	fmt.Println("Email:", data.Person.Email)
	fmt.Println("Mobile:", data.Person.Mobile)
	fmt.Println("Blood Group:", data.Person.BloodGroup)
	fmt.Println("Aadhaar:", data.Person.AadhaarNumber)
}
