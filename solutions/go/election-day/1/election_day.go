package electionday

import "fmt"



func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}

func VoteCount(counter *int) int {
	if counter ==nil{
		return 0
	}
	return int(*counter)
}

func IncrementVoteCount(counter *int, increment int) {
	if counter ==nil{
		return 
	}
	 inc := VoteCount(counter) + increment
	 *counter = inc
}

func NewElectionResult(candidateName string, votes int) *ElectionResult {
	return &ElectionResult{
		Name: candidateName,
		Votes: votes,
	}
}

func DisplayResult(result *ElectionResult) string {
	return  fmt.Sprintf("%s (%v)",result.Name,result.Votes)
}

func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	for name,_:= range results {
		if name==candidate{
		 results[name] = results[name] -1
		}
	}
}
