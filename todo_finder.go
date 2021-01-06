package main

type node struct {
	Name  string
	Type  string
	Nodes []*node
	Todos []*todo
}

type todoFinder struct {
	foundTable map[string]*todo
	foundTree *node
}

func NewTodoFinder() *todoFinder {
	return &todoFinder{
		foundTable: map[string]*todo{},
		foundTree: &node{Name: "root", Type: "dir"},
	}
}

func (tf *todoFinder) AddTodo(t *todo) {
	tf.foundTable[t.Filepath] = t
	tf.foundTree.AddToTree(t.Path(), t)
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
