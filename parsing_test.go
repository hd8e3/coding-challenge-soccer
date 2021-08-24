package main

import (
    "testing"

    "github.com/stretchr/testify/require"
)

// TestParseLine tests the parseLine function.
func TestParseLine(t *testing.T) {
    testCases := []struct{
        message string
        line string
        expectError bool
        expectedErrorString string
        expectedTeamOne string
        expectedTeamTwo string
        expectedTeamOneScore int
        expectedTeamTwoScore int
    }{
        {
            message: "successful parsing",
            line: "Foo Bar 7, Baz 2",
            expectError: false,
            expectedTeamOne: "Foo Bar",
            expectedTeamTwo: "Baz",
            expectedTeamOneScore: 7,
            expectedTeamTwoScore: 2,
        },
        {
            message: "team names end with a number",
            line: "Big Hero 6 7, Fantastic 4 2",
            expectError: false,
            expectedTeamOne: "Big Hero 6",
            expectedTeamTwo: "Fantastic 4",
            expectedTeamOneScore: 7,
            expectedTeamTwoScore: 2,
        },
        {
            message: "pretty confusing teams, but technically still okay",
            line: "One 2, Three 4, Five 6",
            expectError: false,
            expectedTeamOne: "One 2, Three",
            expectedTeamTwo: "Five",
            expectedTeamOneScore: 4,
            expectedTeamTwoScore: 6,
        },
        {
            message: "line doesn't match regexp",
            line: "Foo Bar 7; Baz 2",
            expectError: true,
            expectedErrorString: "regexp did not match line: 'Foo Bar 7; Baz 2'",
        },
    }
    for _, tt := range testCases {
        t.Run(tt.message, func(t *testing.T) {
            team1, score1, team2, score2, err := parseLine(tt.line)
            if tt.expectError {
                require.EqualError(t, err, tt.expectedErrorString)
            } else {
                require.NoError(t, err)
                require.Equal(t, tt.expectedTeamOne, team1)
                require.Equal(t, tt.expectedTeamTwo, team2)
                require.Equal(t, tt.expectedTeamOneScore, score1)
                require.Equal(t, tt.expectedTeamTwoScore, score2)
            }
        })
    }
}
