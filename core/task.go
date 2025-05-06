package core

import (
	"time"

	"github.com/google/uuid"
)

type TaskFn = func(value any) (any, error)

// a unit of work used to generate
// project files incrementally
type Task struct {
	Name    string
	Handler TaskFn
	Pipes   []Task
	Runs    []*TaskRun
}

func NewTask(name string, handler TaskFn) Task {
	return Task{
		Name:    name,
		Handler: handler,
		Pipes:   []Task{},
		Runs:    []*TaskRun{},
	}
}

func (self *Task) Run() (any, error) {
	return self.run(nil)
}

func (self *Task) Pipe(args ...Task) *Task {
	self.Pipes = append(self.Pipes, args...)
	return self
}

func (self *Task) run(value any) (any, error) {
	run := &TaskRun{
		Id:        uuid.NewString(),
		Status:    Running,
		StartedAt: time.Now(),
	}

	self.Runs = append(self.Runs, run)
	res, err := self.Handler(value)

	if err != nil {
		run.Status = Error
		run.Data = err
		return res, err
	}

	for _, task := range self.Pipes {
		res, err = task.run(res)

		if err != nil {
			run.Status = Error
			run.Data = err
			return res, err
		}
	}

	now := time.Now()
	run.EndedAt = &now

	if err != nil {
		run.Status = Error
		run.Data = err
		return res, err
	}

	run.Status = Success
	run.Data = res
	return res, err
}
