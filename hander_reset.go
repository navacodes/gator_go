package main

import (
	"context"
	"fmt"
	"log"
)

func handlerReset(s *State, cmd Command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All users have been deleted")
	return nil
}
