package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/go-getter"
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

// TODO use `getter` to get a local copy of the repo.
//  getter checks if the local files are upto date with the repo, if not updates otherwise continues
func TestNewTodoFinder(t *testing.T) {
	for _, repo := range repos {
		spew.Dump(getter.Get(repo.Dst, repo.Src))
	}
}
