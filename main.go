package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"unicode"
)

// TODO implement dynamic comment token selection, could maybe work similar to entityTracker{}

// TODO check if to-do is just above a function or struct, if so assume the to-do is about the function

const (
	statusDir = "../status-go"
	ignore    = statusDir + "/vendor"
)

var (
	keywords   = []string{"todo", "fixme"}
)

func main() {
	tf := NewTodoFinder()

	err := tf.FindInDir(statusDir)
	if err != nil {
		panic(err)
	}

	spew.Dump(tf.foundTree)
}

func isGoFile(name string) bool {
	if len(name) < 3 {
		return false
	}
	last := name[len(name)-3:]
	return last == ".go"
}

func buildRegexPattern(keywords []string) string {
	kwp := makeRegexKeywords(keywords)
	return fmt.Sprintf("//.*((%s)(.*))", kwp)
}

func makeRegexKeywords(keywords []string) (out string) {
	for i, kw := range keywords {
		for _, r := range kw {
			lower := unicode.ToLower(r)
			upper := unicode.ToUpper(r)

			out += fmt.Sprintf("[%s,%s]", string(lower), string(upper))
		}

		if i+1 < len(keywords) {
			out += "|"
		}
	}

	return
}
