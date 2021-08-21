package main

import (
    "strings"
    "testing"

    "github.com/stretchr/testify/require"
)

// TestSampleOutput is an end-to-end test of the sample input and output.
func TestSampleOutput(t *testing.T) {
    inputLines, err := readAllLinesFromFile("./sample-input.txt")
    require.NoError(t, err)

    outputLines, err := readAllLinesFromFile("./expected-output.txt")
    require.NoError(t, err)

    actualOutput, err := calculateResults(inputLines)
    require.NoError(t, err)
    require.Equal(t, strings.Join(outputLines, "\n") + "\n", actualOutput)
}
