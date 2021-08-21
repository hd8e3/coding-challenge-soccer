package main

import (
    "bufio"
    "flag"
    "io"
    "os"
)

var inputFileName string

func init() {
    flag.StringVar(&inputFileName, "inputFile", "", "name of the input file (if empty, will use stdin)")
}

func readInput() ([]string, error) {
    flag.Parse()
    if inputFileName == "" {
        return readAllLinesFromStdin()
    }
    return readAllLinesFromFile(inputFileName)
}

func readAllLinesFromStdin() ([]string, error) {
    return readAllLinesFromReader(os.Stdin)
}

func readAllLinesFromFile(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    return readAllLinesFromReader(file)
}

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
