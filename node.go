package main

type nodeType int

const (
	DIR nodeType = iota
	FILE
)

type node struct {
	Name  string
	Type  nodeType
	Nodes []*node
	Todos []*todo
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

func (n node) getTypeFromPath(path []string) nodeType {
	if len(path) == 1 {
		return FILE
	}

	return DIR
}

func (n node) String() string {
	return n.toString("", 0)
}

// TODO there is some kind of bug with this. Fix
func (n node) toString(in string, indent int) string {
	idnt := ""
	for i := 0; i < indent; i++ {
		idnt += "  "
	}

	for _, c := range n.Nodes {
		switch c.Type {
		case DIR:
			in += idnt + "- " + c.Name + "\n"
			in = n.toString(in, indent+1)
		case FILE:
			in += idnt + "- " + c.Name + "\n"
		}
	}

	return in
}