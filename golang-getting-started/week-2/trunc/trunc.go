package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter any number you want: ")

	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("ERROR: Failed to read input.")
		os.Exit(0)
	}

	trimmed := strings.TrimSpace(text)
	parsed, err := strconv.ParseFloat(trimmed, 64)

	if err != nil {
		fmt.Println("ERROR: Input was not a valid number.")
		os.Exit(0)
	}

	fmt.Printf("Truncated value: %.0f\n", parsed)
}
