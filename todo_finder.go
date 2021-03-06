package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"unicode"
)

type TodoFinder struct {
	// regex related fields
	entityTracker *entityTracker
	todoRegex     *regexp.Regexp
	lineRegex     *regexp.Regexp

	// results store
	FoundTable []*todo
	foundTree  *node

	// openTodo tracks if subsequent comment lines should be included in the last to-do's description
	openTodo *todo

	keywords []string

	repo *Repo
}

func NewTodoFinder(repo *Repo) (TodoFinder, error) {
	tf := TodoFinder{
		FoundTable: []*todo{},
		foundTree:  &node{Name: "root", Type: DIR},
		keywords:   []string{"todo", "fixme"},
		repo: repo,
	}

	return tf, tf.init()
}

func (tf *TodoFinder) init() (err error) {
	tf.todoRegex, err = regexp.Compile(tf.buildRegexPattern())
	if err != nil {
		return err
	}

	tf.lineRegex, err = regexp.Compile("\n")
	if err != nil {
		return err
	}

	tf.entityTracker, err = NewEntityTracker()
	if err != nil {
		return err
	}

	return nil
}

func (tf *TodoFinder) AddTodo(t *todo) {
	tf.FoundTable = append(tf.FoundTable, t)
	tf.foundTree.AddToTree(t.Path(), t)
}

func (tf *TodoFinder) Find() error {
	if tf.repo == nil {
		return errors.New("repo not set")
	}

	return tf.FindInDir(tf.repo.Dst)
}

func (tf *TodoFinder) FindInDir(dir string) error {
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
			err = tf.FindInDir(filepath)
			if err != nil {
				return err
			}
		}

		if !tf.isValidFile(f.Name()) {
			continue
		}

		file, err := ioutil.ReadFile(filepath)
		if err != nil {
			return err
		}

		lines := tf.lineRegex.Split(string(file), -1)
		for i, l := range lines {
			tf.entityTracker.Track(l)

			results := tf.todoRegex.FindSubmatch([]byte(l))
			if results == nil {
				if len(l) < 3 {
					tf.openTodo = nil
				}

				if tf.openTodo != nil {
					l = strings.TrimSpace(l)
					if l[:2] == "//" {
						tf.openTodo.Description += "\n" + l[2:]
					} else {
						tf.openTodo = nil
					}
				}

				continue
			}
			td := &todo{
				Filepath:          filepath,
				Description:       string(results[1]),
				LineNumber:        i + 1,
				RelatedFuncOrType: tf.entityTracker.Current(),
			}
			tf.AddTodo(td)
			tf.openTodo = td
		}
	}

	return nil
}

func (tf TodoFinder) isValidFile(name string) bool {
	for _, ft := range tf.repo.FileTypes {
		ftl := len(ft) + 1
		if len(name) < ftl {
			return false
		}
		last := name[len(name)-ftl:]
		if last == "."+ft {
			return true
		}
	}
	return false
}

func (tf TodoFinder) buildRegexPattern() string {
	kwp := tf.makeRegexKeywords()
	// TODO add functionality for other types of commenting out
	//  example ';;' or ';' for clojure, '#' and '/* */' blocks
	return fmt.Sprintf("//.*((%s)(.*))", kwp)
}

func (tf TodoFinder) makeRegexKeywords() (out string) {
	for i, kw := range tf.keywords {
		for _, r := range kw {
			lower := unicode.ToLower(r)
			upper := unicode.ToUpper(r)

			out += fmt.Sprintf("[%s,%s]", string(lower), string(upper))
		}

		if i+1 < len(tf.keywords) {
			out += "|"
		}
	}

	return
}
