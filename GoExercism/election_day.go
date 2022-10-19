package main

import "fmt"

type ElectionResult struct {
	// Name of the candidate
	Name string
	// Number of votes the candidate had
	Votes int
}

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}

// VoteCount extracts the number of votes from a counter.

func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	} else {
		return *counter
	}
}

// IncrementVoteCount increments the value in a vote counter.

func IncrementVoteCount(counter *int, increment int) {
	if counter == nil {
		fmt.Println("Don't use nil pointers...")
	} else {
		*counter += increment
	}
}

// NewElectionResult creates a new election result.

func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var ele ElectionResult
	ele.Name = candidateName
	ele.Votes = votes
	return &ele
}

// DisplayResult creates a message with the result to be displayed.

func DisplayResult(result *ElectionResult) string {
	myResult := fmt.Sprintf("%s (%d)", result.Name, result.Votes)
	return myResult
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.

func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] -= 1
}

func main() {
	/*var initialVotes int
	initialVotes = 2

	var counter *int
	counter = NewVoteCounter(initialVotes)
	fmt.Println(*counter == initialVotes)

	var votes int
	votes = 3

	var voteCounter *int
	voteCounter = &votes

	VoteCount(voteCounter)
	fmt.Println(*voteCounter) // => 3

	var nilVoteCounter *int
	fmt.Println(nilVoteCounter)
	a := VoteCount(nilVoteCounter)
	fmt.Println(a) // => 0
	var votes int
	votes = 3

	var voteCounter *int
	voteCounter = &votes

	IncrementVoteCount(voteCounter, 2)

	fmt.Println(votes == 5)        // true
	fmt.Println(*voteCounter == 5) // true

	var result *ElectionResult
	result = NewElectionResult("Peter", 3)

	fmt.Println(result.Name == "Peter") // true
	fmt.Println(result.Votes == 3)      // true
	*/
	var result *ElectionResult
	result = &ElectionResult{
		Name:  "John",
		Votes: 32,
	}

	fmt.Println(DisplayResult(result))
	// => John (32)

	var finalResults = map[string]int{
		"Mary": 10,
		"John": 51,
	}

	DecrementVotesOfCandidate(finalResults, "Mary")

	fmt.Println(finalResults["Mary"])
	// => 9
}
