package main

import (
	"fmt"
	"os"
)

func main() {
	var token = os.Getenv("parcelsApiToken")
	fmt.Println("my token is", token)
}
