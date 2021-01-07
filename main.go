package main

import "fmt"

// TODO implement dynamic comment token selection, could maybe work similar to entityTracker{}

// TODO check if to-do is just above a function or struct, if so assume the to-do is about the function

// TODO output to a file

// TODO add execution params to the running of the application, allowing to set the path, keywords, ignore list, language, etc

const (
	statusDir = "../status-go"
	ignore    = statusDir + "/vendor"
)

func main() {
	tf, err := NewTodoFinder()
	if err != nil {
		panic(err)
	}

	err = tf.FindInDir(statusDir)
	if err != nil {
		panic(err)
	}

	fmt.Println(tf.foundTree.Markdown())
}
