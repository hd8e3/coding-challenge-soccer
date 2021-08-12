package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
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

    if inputFileName == "" {
        allLines = readAllLinesFromStdin()
    } else {
        allLines = readAllLinesFromFile(inputFileName)
    }

    currentMatchDay := 1
    matchDayPlayers := map[string]bool{}
    cumulativeScores := map[string]int{}

    for _, line := range allLines {
        team1, score1, team2, score2 := parseLine(line)

        if contains(matchDayPlayers, team1) || contains(matchDayPlayers, team2) {
            outputMatchDayResults(currentMatchDay, cumulativeScores)
            matchDayPlayers = map[string]bool{}
            currentMatchDay++
        }

        matchDayPlayers[team1] = true
        matchDayPlayers[team2] = true

        if score1 == score2 {
            addToScore(cumulativeScores, team1, 1)
            addToScore(cumulativeScores, team2, 1)
        } else if score1 > score2 {
            addToScore(cumulativeScores, team1, 3)
            addToScore(cumulativeScores, team2, 0)
        } else {
            addToScore(cumulativeScores, team1, 0)
            addToScore(cumulativeScores, team2, 3)
        }
    }

    outputMatchDayResults(currentMatchDay, cumulativeScores)
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

func outputMatchDayResults(currentMatchDay int, matchDayScores map[string]int) {
    fmt.Printf("Matchday %v\n", currentMatchDay)

    sortable := NewSortable(matchDayScores)
    sort.Sort(sortable)

    for i := 0; i < len(sortable.teams) && i < 3; i++ {
        fmt.Printf("%v, %v pt%v\n", sortable.teams[i], sortable.scores[i], pluralize(sortable.scores[i]))
    }
    fmt.Println()
}

func pluralize(score int) string {
    if score == 1 {
        return ""
    }
    return "s"
}

func parseLine(line string) (string, int, string, int) {
    matches := lineRegexp.FindStringSubmatch(line)
    if matches == nil || len(matches) != 5 {
        log.Fatal("regexp did not match line")
    }
    score1, err1 := strconv.Atoi(matches[2])
    score2, err2 := strconv.Atoi(matches[4])
    if err1 != nil || err2 != nil {
        log.Fatal("unable to convert score to int")
    }

    return matches[1], score1, matches[3], score2
}

func readAllLinesFromStdin() []string {
    var ret []string
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        ret = append(ret, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return ret
}

func readAllLinesFromFile(filename string) []string {
    var ret []string
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // TODO DRY
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        ret = append(ret, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return ret
}
