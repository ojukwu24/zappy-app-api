package main

import "fmt"

// Run is responsible for instantiating
// the startup of the application
func Run() error {
	fmt.Println("Starting server...")
	return nil
}

func main() {
	fmt.Println("Zappy API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
