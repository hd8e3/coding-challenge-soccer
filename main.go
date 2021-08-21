package main

import (
    "fmt"
)

// main is the main entrypoint into the program.
func main() {
    inputLines, err := readInput()
    if err != nil {
        fmt.Printf("Error reading input: %v", err)
        return
    }

    resultString, err := calculateResults(inputLines)
    if err != nil {
        fmt.Printf("Error in result calculation: %v", err)
    } else {
        fmt.Print(resultString)
    }
}
