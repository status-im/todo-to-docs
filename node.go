package main

type node struct {
	Name  string
	Type  string
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

func (n node) getTypeFromPath(path []string) string {
	if len(path) == 1 {
		return "file"
	}

	return "dir"
}
