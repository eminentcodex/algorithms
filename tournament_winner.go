package main

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.
    var leadingTeam, currentTeam string
    var leadingScore, currentScore int
    var winners = make(map[string]int)

    for i, c := range competitions {
        if results[i] == HOME_TEAM_WON {
            if _, ok := winners[c[0]]; !ok {
                winners[c[0]] = 0
            }
            winners[c[0]] = winners[c[0]] + 3
            currentScore = winners[c[0]]
            currentTeam = c[0]
        } else {
            if _, ok := winners[c[1]]; !ok {
                winners[c[1]] = 0
            }
            winners[c[1]] = winners[c[1]] + 3
            currentScore = winners[c[1]]
            currentTeam = c[1]
        }

        if currentScore > leadingScore {
            leadingScore = currentScore
            leadingTeam = currentTeam
        }
    }
    
	return leadingTeam
}
