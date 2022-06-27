package selectingTheRightJobs

import (
	"sort"
)

func EarliestFirst(jobs []Job) (result []string) {
	var maxTime int
	for _, v := range jobs {
		if v.end > maxTime {
			maxTime = v.end
		}
	}

	for currentTime := 0; currentTime <= maxTime; currentTime++ {
		for i, v := range jobs {
			if v.start == currentTime {
				result = append(result, v.name)
				currentTime = v.end - 1
				jobs = removeJob(jobs, i)
				break
			}
		}
	}

	return
}

func EarliestFirstV2(jobs []Job) (result []string) {
	var time int
	sortByStartTime(jobs)
	for _, job := range jobs {
		if job.start >= time {
			time = job.end
			result = append(result, job.name)
		}
	}
	return
}

func ShortestFirst(jobs []Job) (result []string) {
	var addedJobs []Job
	jobs = sortByDuration(jobs)
	for _, v := range jobs {
		if !hasOverlap(v, addedJobs) {
			addedJobs = append(addedJobs, v)
		}
	}
	sortByStartTime(addedJobs)
	result = mapJobToResult(addedJobs)
	return
}

func OptimalScheduling(jobs []Job) (result []string) {
	sortByEndTime(jobs)
	var time int
	for _, job := range jobs {
		if job.start >= time {
			result = append(result, job.name)
			time = job.end
		}
	}
	return
}

func sortByEndTime(jobs []Job) []Job {
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].end < jobs[j].end
	})
	return jobs
}

func sortByStartTime(jobs []Job) []Job {
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].start < jobs[j].start
	})
	return jobs
}

func sortByDuration(jobs []Job) []Job {
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].Duration() < jobs[j].Duration()
	})
	return jobs
}

func hasOverlap(job Job, occupied []Job) bool {
	for _, v := range occupied {
		if job.start > v.start && job.start < v.end ||
			job.end > v.start && job.end < v.end ||
			v.start > job.start && v.start < job.end ||
			v.end > job.start && v.end < job.end {
			return true
		}
	}
	return false
}

func mapJobToResult(data []Job) []string {
	mapped := make([]string, len(data))
	for i, e := range data {
		mapped[i] = e.name
	}
	return mapped
}

func removeJob(s []Job, i int) []Job {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
