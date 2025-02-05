package main

import (
    "strings"
    "testing"

    "github.com/stretchr/testify/require"
)

// TestCalculateResults tests calculateResults with the sample input and output.
func TestCalculateResults(t *testing.T) {
    inputLines, err := readAllLinesFromFile("./sample-input.txt")
    require.NoError(t, err)

    outputLines, err := readAllLinesFromFile("./expected-output.txt")
    require.NoError(t, err)

    actualOutput, err := calculateResults(inputLines)
    require.NoError(t, err)
    require.Equal(t, strings.Join(outputLines, "\n") + "\n", actualOutput)
}

// TestCalculateResultsWithError tests that we correctly error out when we can't parse a line.
func TestCalculateResultsWithError(t *testing.T) {
    _, err := calculateResults([]string{
        "Foo Bar",
        "Baz",
    })
    require.Error(t, err)
}
