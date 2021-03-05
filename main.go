package main

import (
	"fmt"
	"github.com/hashicorp/go-getter"
)

// TODO add support for `status-react`

// TODO implement dynamic comment token selection, could maybe work similar to entityTracker{}

// TODO check if to-do is just above a function or struct, if so assume the to-do is about the function

// TODO output to a file

// TODO add execution params to the running of the application, allowing to set the path, keywords, ignore list, language, etc

const (
	statusDir = "../status-go"
	ignore    = statusDir + "/vendor"
)

type Repo struct {
	Dst string
	Src string
	FileTypes []string // To be used with (tf TodoFinder) isValidFile
}

var (
	repos = []Repo{
		{"./.dst/status-go", "github.com/status-im/status-go", []string{"go"}},
		{"./.dst/status-react", "github.com/status-im/status-react", []string{"cljs"}},
	}
)

func main() {
	repo := &Repo{
		Dst:  ".dst/todos-to-docs",
		Src: "github.com/status-im/todo-to-docs",
		FileTypes: []string{"go"},
	}

	err := getter.Get(repo.Dst, repo.Src)
	if err != nil {
		panic(err)
	}

	tf, err := NewTodoFinder(repo)
	if err != nil {
		panic(err)
	}

	err = tf.Find()
	if err != nil {
		panic(err)
	}

	fmt.Println(tf.foundTree)
}
