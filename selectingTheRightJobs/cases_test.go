package selectingTheRightJobs

import (
	"math/rand"
	"time"
)

var testCasesOptimal = []struct {
	s1          []Job
	want        []string
	expectError bool
}{
	{
		append([]Job(nil), shuffledJobs...),
		[]string{
			"Discrete Mathematics",
			"Halting State",
			"Programming Challenges",
			"Calculated Bets",
		},
		false,
	},
}

var testCasesEarliest = []struct {
	s1          []Job
	want        []string
	expectError bool
}{
	{
		append([]Job(nil), shuffledJobs...),
		[]string{
			"The President's Algorist",
			"Steiner's Tree",
			"Programming Challenges",
			"Calculated Bets",
		},
		false,
	},
}

var testCasesShortest = []struct {
	s1          []Job
	want        []string
	expectError bool
}{
	{
		append([]Job(nil), shuffledJobs...),
		[]string{
			"Discrete Mathematics",
			"Steiner's Tree",
			"Programming Challenges",
			"Calculated Bets",
		},
		false,
	},
}

var jobs = []Job{
	{
		name:  "The President's Algorist",
		start: 0,
		end:   10,
	},
	{
		name:  "Discrete Mathematics",
		start: 2,
		end:   8,
	},
	{
		name:  "Tarjan of the Jungle",
		start: 4,
		end:   13,
	},
	{
		name:  "Halting State",
		start: 9,
		end:   14,
	},
	{
		name:  "Steiner's Tree",
		start: 12,
		end:   16,
	},
	{
		name:  "The Four Volume Problem",
		start: 15,
		end:   30,
	},
	{
		name:  "Programming Challenges",
		start: 17,
		end:   28,
	},
	{
		name:  "Process Terminated",
		start: 25,
		end:   35,
	},
	{
		name:  "Calculated Bets",
		start: 29,
		end:   36,
	},
}

var shuffledJobs = GetShuffledJobs()

func GetShuffledJobs() []Job {
	result := append([]Job(nil), jobs...)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })
	return result
}
