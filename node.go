package main

import "strings"

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

func (n node) Markdown() string {
	s := n.String()
	return strings.ReplaceAll(s, "\n", "<br/>\n")
}

func (n node) String() string {
	return n.toString(0)
}

func (n node) toString(indent int) string {
	out := ""
	baseIdnt := "  "

	idnt := ""
	for i := 0; i < indent; i++ {
		idnt += baseIdnt
	}

	for _, c := range n.Nodes {
		switch c.Type {
		case DIR:
			out += idnt + "- 📁 " + c.Name + "\n"
			out += c.toString(indent+1)
		case FILE:
			out += idnt + "- 📃 " + c.Name + "\n"
			for _, t := range c.Todos {
				todoIdnt := idnt+baseIdnt+baseIdnt
				ts := strings.ReplaceAll(t.String(), "\n", "\n"+todoIdnt)
				out += idnt + baseIdnt + "- ⬜ Todo:\n" + todoIdnt + ts + "\n"
			}
		}
	}

	return out
}