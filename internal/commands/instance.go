package commands

import (
	"fmt"
	"github.com/EdmundMartin/gstalk/internal/job"
	"github.com/EdmundMartin/gstalk/internal/structures"
	"github.com/EdmundMartin/gstalk/internal/tube"
)

type Instance struct {
	Tubes map[string]*tube.Tube
}


func NewInstance() *Instance {
	return &Instance{Tubes: map[string]*tube.Tube{}}
}


func (i *Instance) PutJob(tubeName string, job *job.Job) error {
	t, ok := i.Tubes[tubeName]
	if !ok {
		i.Tubes[tubeName] = tube.NewTube(tubeName)
		t, _ = i.Tubes[tubeName]
	}
	t.Insert(job)
	return nil
}

func (i *Instance) GetJob(tubeNames structures.StringSet) (*job.Job, error){
	var mostUrgent int
	var targetTube *tube.Tube
	for name, currentTube := range i.Tubes {
		if tubeNames.Contains(name) {
			fmt.Println(name)
			t, pri := isHigherPriority(targetTube, mostUrgent, currentTube)
			if t != nil {
				targetTube = t
				mostUrgent = pri
			}
		}
	}
	if targetTube == nil {
		return nil, fmt.Errorf("no job ready")
	}
	j, err := targetTube.Remove()
	if err != nil {
		return nil, err
	}
	return j, nil
}


func isHigherPriority(selected *tube.Tube, urgency int, tube *tube.Tube) (*tube.Tube, int) {
	if selected == nil {
		return tube, tube.Peek().Priority()
	}
	nextJob := tube.Peek()
	if nextJob != nil && nextJob.Priority() < urgency {
		return tube, nextJob.Priority()
	}
	return nil, -1
}


func (i *Instance) TotalJobsReady() {

}