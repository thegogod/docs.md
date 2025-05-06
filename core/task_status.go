package core

type TaskStatus string

const (
	Idle    TaskStatus = "idle"
	Running TaskStatus = "running"
	Success TaskStatus = "success"
	Error   TaskStatus = "error"
)
