package core

type TaskOutput struct {
	Value any
	Error error
	Run   TaskRun
}
