package tube

import (
	"fmt"
	"github.com/EdmundMartin/gstalk/internal/heap"
	"github.com/EdmundMartin/gstalk/internal/job"
	"os"
)


type Tube struct {
	Name string
	jobs *heap.MinHeap
	delayed *heap.MinHeap
	reserved *heap.MinHeap
}

func NewTube(name string) *Tube {
	return &Tube{
		Name: name,
		jobs: heap.NewMinHeap(),
		delayed: heap.NewMinHeap(),
		reserved: heap.NewMinHeap(),
	}
}


func (t *Tube) Insert(j *job.Job) {
	t.jobs.Insert(j)
}

func (t *Tube) HighestPriority() int {
	j := t.jobs.Peek()
	if j == nil {
		return -1
	}
	return j.Priority()
}

func (t *Tube) Remove() (*job.Job, error) {
	j := t.jobs.Remove()
	if j == nil {
		return nil, fmt.Errorf("no jobs in tube")
	}
	return j.(*job.Job), nil
}

func (t *Tube) Size() int {
	return t.jobs.Size()
}

func (t *Tube) Peek() *job.Job {
	j := t.jobs.Peek()
	return j.(*job.Job)
}


func idxToSectionName(idx int) string {
	switch idx {
	case 0:
		return "__JOBS__\n"
	case 1:
		return "__DELAYED__\n"
	case 2:
		return "__RESERVED__\n"
	default:
		return ""
	}
}

func (t *Tube) SyncToFile(filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	for idx, section := range []*heap.MinHeap{t.jobs, t.delayed, t.reserved} {
		_, err = f.WriteString(idxToSectionName(idx))
		if err != nil {
			return err
		}
		for _, j := range section.Contents() {
			currentJob := j.(*job.Job)
			_, err := f.WriteString(currentJob.String() + "\n")
			if err != nil {
				return err
			}
		}
	}
	return nil
}