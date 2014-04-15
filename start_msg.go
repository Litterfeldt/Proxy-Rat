package main

import "fmt"

func printStartupMessage() {
	fmt.Println()
	fmt.Println("Proxy rat is starting up!\n")

	fmt.Printf("port:        %s\n", *port)

	fmt.Println()
}
