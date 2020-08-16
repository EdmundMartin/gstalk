package job

import "fmt"

type Job struct {
	content string
	priority int
	status Status
	ttr int
	attempts int
}

func NewJob(content string, priority int, ttr int) *Job {
	return &Job{
		content:  content,
		priority: priority,
		ttr:      ttr,
		status:   Ready,
		attempts: 0,
	}
}

func (j *Job) Release() {
	j.status = Ready
}

func (j *Job) Reserve() {
	j.status = Reserved
}

func (j *Job) Priority() int {
	return j.priority
}


func (j *Job) String() string {
	return fmt.Sprintf("%s %d %d %s %d", j.content, j.priority, j.ttr, j.status, j.attempts)
}