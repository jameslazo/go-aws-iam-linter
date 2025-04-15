// main.go
package main

import (
	"fmt"
	"os"
	_ "encoding/json"
	// "aws-iam-linter/internal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: aws-iam-linter <iam.json>")
		return
	}

	filePath := os.Args[1]
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("File read failure: %v\n", err)
	} else {
			fmt.Println(string(data))
	}
}
