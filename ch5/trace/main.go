package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(5 * time.Second) // simulate slow operation by sleeping
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("entering %s", msg)
	return func() { log.Printf("exit %s, cost %s", msg, time.Since(start)) }
}

func main() {
	bigSlowOperation()
}
