package main

import (
	"fmt"
	"github.com/navacodes/gator_go/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		return
	}
	fmt.Printf("Initial config: %+v\n", cfg)
	// Step 2: Set the current user to your name
	err = cfg.SetUser("David")
	if err != nil {
		fmt.Printf("Error setting user: %v\n", err)
		return
	}
	// Step 3: Read the config again and print it
	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Error reading config again: %v\n", err)
		return
	}
	fmt.Printf("Updated config: %+v\n", cfg)
}
