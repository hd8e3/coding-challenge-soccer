package main

// Sortable implements the sort.Interface interface so that we can sort teams by their scores.
type Sortable struct {
    teams []string
    scores []int
}

// NewSortable constructs a new Sortable.
func NewSortable(theMap map[string]int) *Sortable {
    var teams []string
    var scores []int

    for k, v := range theMap {
        teams = append(teams, k)
        scores = append(scores, v)
    }

    return &Sortable{teams, scores}
}

// Len is the number of elements in the collection.
func (s *Sortable) Len() int {
    return len(s.teams)
}

// Less reports whether the element with index i must sort before the element with index j.
func (s *Sortable) Less(i, j int) bool {
    if s.scores[i] == s.scores[j] {
        // If scores are identical, compare team names lexicographically
        return s.teams[i] < s.teams[j]
    }
    return s.scores[i] > s.scores[j]
}

// Swap swaps the elements with indexes i and j.
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
