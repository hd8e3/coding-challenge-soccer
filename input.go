package main

import (
    "bufio"
    "flag"
    "io"
    "os"
)

// inputFileName is the file name containing input.
var inputFileName string

// init initializes the inputFileName from the command-line flag.
func init() {
    flag.StringVar(&inputFileName, "inputFile", "", "name of the input file (if empty, will use stdin)")
}

// readInput returns a slice of all the input lines.
func readInput() ([]string, error) {
    flag.Parse()
    if inputFileName == "" {
        return readAllLinesFromStdin()
    }
    return readAllLinesFromFile(inputFileName)
}

// readAllLinesFromStdin returns a slice of all the input lines from stdin.
func readAllLinesFromStdin() ([]string, error) {
    return readAllLinesFromReader(os.Stdin)
}

// readAllLinesFromStdin returns a slice of all the input lines from the file identified by filename.
func readAllLinesFromFile(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    return readAllLinesFromReader(file)
}

// readAllLinesFromReader returns a slice of all the input lines read from the io.Reader r.
func readAllLinesFromReader(r io.Reader) ([]string, error) {
    var ret []string
    scanner := bufio.NewScanner(r)
    for scanner.Scan() {
        ret = append(ret, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return ret, nil
}
