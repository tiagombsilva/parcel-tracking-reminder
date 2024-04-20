package main

import (
	"fmt"
	"parcelsApi/handler"
	"strings"
)

func main() {
	response := handler.GetParcel("6623e161af2ead75e45f42b4")
	fmt.Println(strings.TrimSpace(response.UUID))
}
