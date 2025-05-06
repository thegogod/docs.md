package core

import "time"

// represents data regarding
// some tasks execution
type TaskRun struct {
	Id        string     `json:"id"`
	Status    TaskStatus `json:"status"`
	Data      any        `json:"data,omitempty"`
	StartedAt time.Time  `json:"started_at"`
	EndedAt   *time.Time `json:"ended_at,omitempty"`
}
