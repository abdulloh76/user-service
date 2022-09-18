package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("DYNAMODB_PORT", os.Getenv("DYNAMODB_PORT"))
	fmt.Println("PORT", os.Getenv("PORT"))
}
