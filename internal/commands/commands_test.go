package commands

import (
	"github.com/EdmundMartin/gstalk/internal/job"
	"github.com/EdmundMartin/gstalk/internal/structures"
	"testing"
)

func TestInstance_PutJob(t *testing.T) {
	instance := NewInstance()
	_ = instance.PutJob("default", job.NewJob("test", 100, 20))
	res, _ := instance.GetJob(structures.HashSetFromSlice([]string{"default"}))
	result := res.String()
	if result != "test 100 20 ready 0" {
		t.Error("did not receive expected job back from queue")
	}
}

func TestInstance_PutManyJobs(t *testing.T) {
	inst := NewInstance()
	allJobs := []*job.Job{ job.NewJob("test", 25, 20), job.NewJob("highest", 20, 20)}
	for _, j := range allJobs {
		inst.PutJob("default", j)
	}
	res, _ := inst.GetJob(structures.HashSetFromSlice([]string{"default"}))
	if res.String() != "highest 20 20 ready 0" {
		t.Error("did not get the most urgent job")
	}
}


func TestInstance_PutManyJobsManyTubes(t *testing.T) {
	inst := NewInstance()
	defaultJobs := []*job.Job{job.NewJob("test", 25, 20), job.NewJob("highest", 20, 20)}
	otherJobs := []*job.Job{job.NewJob("otherTest", 18, 20), job.NewJob("otherHighest", 10, 10)}
	for _, j := range defaultJobs {
		inst.PutJob("default", j)
	}
	for _, j := range otherJobs {
		inst.PutJob("other", j)
	}
	res, _ := inst.GetJob(structures.HashSetFromSlice([]string{"default", "other"}))
	if res.String() != "otherHighest 10 10 ready 0" {
		t.Errorf("did not get most urgent job from tube: %s", res.String())
	}
}