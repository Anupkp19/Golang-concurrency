package main

import (
	"fmt"
	"log"
	"strconv"

)

func killServer(pidFile string) error {
	// TODO

	// Simulate kill
	pid, err := strconv.Atoi(pidFile)
	if err != nil {
		return err
	}
	fmt.Printf("killing server with pid=%d\n", pid)
	return nil
}

func main() {
	if err := killServer("serverpid"); err != nil {
		log.Fatalf("error: %s", err)
	}
}
