package main

import (
	"io/ioutil"
	"regexp"

	"github.com/davecgh/go-spew/spew"
)

const (
	statusDir = "../status-go"
	ignore = statusDir + "/vendor"
)

// TODO implement dynamic comment token selection
// TODO implement dynamic todo phrase search generation
//  ie generate regex string from a slice of keywords

var (
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
	r, err := regexp.Compile("//.*(([t,T][o,O][d,D][o,O]|[f,F][i,I][x,X][m,M][e,E])(.*))[\n,\r,\n\r,\r\n]")
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
			spew.Dump(rst)
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
