package selectingTheRightJobs

import "fmt"

type Job struct {
	name  string
	start int
	end   int
}

// TimeSlots key is start time and value is duration
type TimeSlots map[int]int

func (job Job) String() string {
	return fmt.Sprintf("\"%s\"\n", job.name)
}

func (job Job) Duration() int {
	return job.end - job.start
}
