package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

// TODO implement dynamic comment token selection

// TODO With line number find if todo is in a function or struct record the name of the function / struct in `todo{}`

const (
	statusDir = "../status-go"
	ignore = statusDir + "/vendor"
)

var (
	keywords = []string{"todo", "fixme"}
	found []todo

	// todoMode tracks if subsequent comment lines should be included in the last to-do's description
	todoMode = false

)

type todo struct {
	Filepath string
	Description string
	LineNumber int
}

func main() {
	err := processFilesInDir(statusDir)
	if err != nil {
		panic(err)
	}

	spew.Dump(found)
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

		if !isGoFile(f.Name()){
			continue
		}

		file, err := ioutil.ReadFile(filepath)
		if err != nil {
			return err
		}

		lines := lineSplitter.Split(string(file), -1)
		for i, l := range lines {

			results := r.FindSubmatch([]byte(l))
			if results == nil {
				if len(l) < 3 {
					todoMode = false
				}

				if todoMode {
					l = strings.TrimSpace(l)
					if l[:2] == "//" {
						found[len(found)-1].Description += "\n" + l[2:]
					} else {
						todoMode = false
					}
				}

				continue
			}
			found = append(found, todo{filepath, string(results[1]), i+1})
			todoMode = true

			// TODO Working on getting subsequent lines of a todo
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
