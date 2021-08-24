package main

import (
    "fmt"
    "regexp"
    "strconv"
)

// lineRegexp is a regular expression representing the format of each input line.
var lineRegexp = regexp.MustCompile(`(.+) ([\d]+), (.+) ([\d]+)`)

// parseLine parses a single line. It returns team1, team1's score, team2, and team2's score, or an
// error upon parsing failure.
func parseLine(line string) (string, int, string, int, error) {
    matches := lineRegexp.FindStringSubmatch(line)
    if matches == nil || len(matches) != 5 {
        return "", 0, "", 0, fmt.Errorf("regexp did not match line: '%v'", line)
    }

    // can safely ignore the errors here because a string matching [\d]+ will always be convertible to an int
    score1, _ := strconv.Atoi(matches[2])
    score2, _ := strconv.Atoi(matches[4])

    return matches[1], score1, matches[3], score2, nil
}
