package main

import (
    "fmt"
    "sort"
)

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
