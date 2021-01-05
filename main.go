package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

// TODO implement dynamic comment token selection

// TODO get line number of todo
//  can be found by splitting the file by line break and regex find on each line.

// TODO add line number to `todo` struct

// TODO With line number find if todo is in a function or struct record the name of the function / struct in `todo{}`

const (
	statusDir = "../status-go"
	ignore = statusDir + "/vendor"
)

var (
	keywords = []string{"todo", "fixme"}
	found []todo
)

type todo struct {
	Filepath string
	Description string
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

		results := r.FindAllSubmatch(file, -1)
		if results == nil {
			continue
		}

		for _, rst := range results {
			found = append(found, todo{filepath, string(rst[1])})
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
	return fmt.Sprintf("//.*((%s)(.*))[\n,\n,\n\n,\n]", kwp)
}

func makeRegexKeywords(keywords []string) string {
	var out string

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

	return out
}
