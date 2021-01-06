package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

type node struct {
	Name  string
	Type  string
	Nodes []*node
	Todos []*todo
}

type TodoFinder struct {
	FoundTable []*todo
	foundTree  *node

	// openTodo tracks if subsequent comment lines should be included in the last to-do's description
	openTodo *todo
}

func NewTodoFinder() TodoFinder {
	return TodoFinder{
		FoundTable: []*todo{},
		foundTree:  &node{Name: "root", Type: "dir"},
	}
}

func (tf *TodoFinder) AddTodo(t *todo) {
	tf.FoundTable = append(tf.FoundTable, t)
	tf.foundTree.AddToTree(t.Path(), t)
}

func (tf *TodoFinder) FindInDir(dir string) error {
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
				RelatedFuncOrType: et.Current(),
			}
			tf.AddTodo(td)
			tf.openTodo = td
		}
	}

	return nil
}

func (n *node) AddToTree(path []string, t *todo) {
	if len(path) == 0 {
		n.Todos = append(n.Todos, t)
		return
	}

	var nn *node
	for _, cn := range n.Nodes {
		if cn.Name == path[0] {
			nn = cn
		}
	}

	if nn == nil {
		nn = &node{
			Name: path[0],
			Type: n.getTypeFromPath(path),
		}

		n.Nodes = append(n.Nodes, nn)
	}

	nn.AddToTree(path[1:], t)
}

func (n node) getTypeFromPath(path []string) string {
	if len(path) == 1 {
		return "file"
	}

	return "dir"
}

func trimPath(path []string) []string {
	ignoreList := []string{"..", "status-go"}
	if path[0] == ignoreList[0] && path[1] == ignoreList[1] {
		path = path[2:]
	}

	return path
}
