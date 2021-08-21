package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "os"
    "regexp"
    "sort"
    "strconv"
)

type Sortable struct {
    teams []string
    scores []int
}

func NewSortable(theMap map[string]int) *Sortable {
    var teams []string
    var scores []int

    for k, v := range theMap {
        teams = append(teams, k)
        scores = append(scores, v)
    }

    return &Sortable{teams, scores}
}

func (s *Sortable) Len() int {
    return len(s.teams)
}

func (s *Sortable) Less(i, j int) bool {
    if s.scores[i] == s.scores[j] {
        // If scores are identical, compare team names lexicographically
        return s.teams[i] < s.teams[j]
    }
    return s.scores[i] > s.scores[j]
}

func (s *Sortable) Swap(i, j int) {
    team1 := s.teams[i]
    team2 := s.teams[j]
    score1 := s.scores[i]
    score2 := s.scores[j]

    s.teams[i] = team2
    s.teams[j] = team1
    s.scores[i] = score2
    s.scores[j] = score1
}

var inputFileName string
var lineRegexp = regexp.MustCompile(`(.+) ([\d]+), (.+) ([\d]+)`)

func init() {
    flag.StringVar(&inputFileName, "inputFile", "", "name of the input file (if empty, will use stdin)")
}

func main() {
    flag.Parse()

    var allLines []string
    var err error

    if inputFileName == "" {
        allLines, err = readAllLinesFromStdin()
        if err != nil {
            fmt.Printf("Error reading from stdin: %v", err)
            return
        }
    } else {
        allLines, err = readAllLinesFromFile(inputFileName)
        if err != nil {
            fmt.Printf("Error reading from file '%v': %v", inputFileName, err)
            return
        }
    }

    resultString, err := calculateResults(allLines)
    if err != nil {
        fmt.Printf("Encountered an error in result calculation: %v", err)
    } else {
        fmt.Print(resultString)
    }
}

func calculateResults(inputLines []string) (string, error) {
    ret := "" // The string we will eventually return

    currentMatchDay := 1 // State: current match day being processed
    matchDayPlayers := map[string]bool{} // State: for the current match day, keep track of which players we've seen already
    cumulativeScores := map[string]int{} // State: cumulative scores so far, across match days

    for _, line := range inputLines {
        team1, score1, team2, score2, err := parseLine(line)
        if err != nil {
            return "", err
        }

        if contains(matchDayPlayers, team1) || contains(matchDayPlayers, team2) {
            // Repeat player indicates end of previous match day.
            ret += matchDayResults(currentMatchDay, cumulativeScores) + "\n" // Append match day results to output
            matchDayPlayers = map[string]bool{} // Reset players we've seen so far for current match day
            currentMatchDay++
        }

        // Mark team1 and team2 as having been seen for current match day
        matchDayPlayers[team1] = true
        matchDayPlayers[team2] = true

        if score1 == score2 {
            addToScore(cumulativeScores, team1, 1)
            addToScore(cumulativeScores, team2, 1)
        } else if score1 > score2 {
            addToScore(cumulativeScores, team1, 3)
            addToScore(cumulativeScores, team2, 0) // Add zero to ensure the team is present.
        } else {
            addToScore(cumulativeScores, team1, 0) // Add zero to ensure the team is present.
            addToScore(cumulativeScores, team2, 3)
        }
    }

    ret += matchDayResults(currentMatchDay, cumulativeScores) // Nothing left to process. Append final match day results to output.
    return ret, nil
}

func addToScore(scores map[string]int, team string, scoreToAdd int) {
    v, ok := scores[team]
    if !ok {
        v = 0
    }
    scores[team] = v + scoreToAdd
}

func contains(theMap map[string]bool, team string) bool {
    _, ok := theMap[team]
    return ok
}

func matchDayResults(currentMatchDay int, matchDayScores map[string]int) string {
    ret := ""
    ret += fmt.Sprintf("Matchday %v\n", currentMatchDay)

    sortable := NewSortable(matchDayScores)
    sort.Sort(sortable)

    for i := 0; i < len(sortable.teams) && i < 3; i++ {
        ret += fmt.Sprintf("%v, %v pt%v\n", sortable.teams[i], sortable.scores[i], pluralize(sortable.scores[i]))
    }
    return ret
}

func pluralize(score int) string {
    if score == 1 {
        return ""
    }
    return "s"
}

func parseLine(line string) (string, int, string, int, error) {
    matches := lineRegexp.FindStringSubmatch(line)
    if matches == nil || len(matches) != 5 {
        return "", 0, "", 0, fmt.Errorf("regexp did not match line: %v", line)
    }
    score1, err1 := strconv.Atoi(matches[2])
    if err1 != nil {
        return "", 0, "", 0, fmt.Errorf("unable to convert score to int: %v", matches[2])
    }
    score2, err2 := strconv.Atoi(matches[4])
    if err2 != nil {
        return "", 0, "", 0, fmt.Errorf("unable to convert score to int: %v", matches[4])
    }

    return matches[1], score1, matches[3], score2, nil
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
