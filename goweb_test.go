package main

import (
	"os"
	"testing"
)

var (
	cm   CommandRegister
	c    ControllerRegister
	s    ServiceRegister
	m    ModelRegister
	mw   interface{}
	args = []string{
		"show:commands",
	}
)

func TestStart(t *testing.T) {
	pwd, _ := os.Getwd()
	if err := os.Setenv("base_path", pwd); err != nil {
		ProcessError(err)
	}

	Start(args, cm, c, s, mw, m)
}
