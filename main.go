package main

import (
	"github.com/davecgh/go-spew/spew"
)

// TODO implement dynamic comment token selection, could maybe work similar to entityTracker{}

// TODO check if to-do is just above a function or struct, if so assume the to-do is about the function

const (
	statusDir = "../status-go"
	ignore    = statusDir + "/vendor"
)

func main() {
	tf := NewTodoFinder()
	err := tf.Init()
	if err != nil {
		panic(err)
	}

	err = tf.FindInDir(statusDir)
	if err != nil {
		panic(err)
	}

	spew.Dump(tf.FoundTable)
}
