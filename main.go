package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
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

	// openTodo tracks if subsequent comment lines should be included in the last to-do's description
	openTodo *todo
)

type todo struct {
	Filepath          string
	Description       string
	LineNumber        int
	RelatedFuncOrType string
	filePathSlice	  []string
}

func (t *todo) Path() []string {
	if t.filePathSlice != nil {
		return t.filePathSlice
	}

	sp := strings.Split(t.Filepath, "/")
	sp = trimPath(sp)
	t.filePathSlice = sp

	return t.filePathSlice
}

func main() {
	err := processFilesInDir(statusDir)
	if err != nil {
		panic(err)
	}
}

func processFilesInDir(dir string) error {
	r, err := regexp.Compile(buildRegexPattern(keywords))
	if err != nil {
		return err
	}

	lineSplitter, err := regexp.Compile("\n")
	if err != nil {
		return err
	}

	et, err := NewEntityTracker()
	if err != nil {
		return err
	}

	tf := NewTodoFinder()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		filepath := dir + "/" + f.Name()

		if filepath == ignore {
			continue
		}

		if f.IsDir() {
			err = processFilesInDir(filepath)
			if err != nil {
				return err
			}
		}

		if !isGoFile(f.Name()) {
			continue
		}

		file, err := ioutil.ReadFile(filepath)
		if err != nil {
			return err
		}

		lines := lineSplitter.Split(string(file), -1)
		for i, l := range lines {
			et.Track(l)

			results := r.FindSubmatch([]byte(l))
			if results == nil {
				if len(l) < 3 {
					openTodo = nil
				}

				if openTodo != nil {
					l = strings.TrimSpace(l)
					if l[:2] == "//" {
						openTodo.Description += "\n" + l[2:]
					} else {
						openTodo = nil
					}
				}

				continue
			}
			td := &todo{
				Filepath:          filepath,
				Description:       string(results[1]),
				LineNumber:        i + 1,
				RelatedFuncOrType: et.Current(),
			}
			tf.AddTodo(td)
			openTodo = td
		}
	}

	return nil
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
