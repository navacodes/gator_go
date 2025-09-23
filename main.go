package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/navacodes/gator_go/internal/config"
	"github.com/navacodes/gator_go/internal/database"
)

func main() {
	cfg, err := config.Read()
	dbURL := cfg.DBURL
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	programState := &State{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := CommandRegistry{
		registeredCommands: make(map[string]func(*State, Command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	if len(os.Args) < 2 {
		log.Fatal("Usage: gator <command> [args...]")
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, Command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
