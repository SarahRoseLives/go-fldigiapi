package main

import (
	"fldigi/fldigiapi" // Correct the import path
	"fmt"
	"log"
)

func main() {
	// Create a new XML-RPC client
	client := fldigiapi.NewClient("http://localhost:7362") // Correct the package reference

	// Call the fldigi.config_dir method
	response, err := client.Call("main.set_frequency", nil)
	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Println("Response:", response)

}
