package main

type command struct {
	Name   string
	Usage  string
	Action func(*application) bool
}

func newCommand(name string, usage string, action func(*application) bool) command {
	return command{
		Name:   name,
		Usage:  usage,
		Action: action,
	}
}
