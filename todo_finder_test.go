package main

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestNode_AddToTree(t *testing.T) {
	n := &node{Name: "root", Type: DIR}
	td := &todo{
		Filepath:          "../status-go/admin/sales/donkeys/pokemon/gif.go",
		Description:       "this looks borken",
		LineNumber:        19224,
		RelatedFuncOrType: "",
	}

	n.AddToTree(td.Path(), td)

	spew.Dump(n)
}
