package main

import "strings"

type todo struct {
	Filepath          string
	Description       string
	LineNumber        int
	RelatedFuncOrType string
	filePathSlice	  []string
}

func (t *todo) Path() []string {
	if t.filePathSlice != nil {
		return t.filePathSlice
	}

	sp := strings.Split(t.Filepath, "/")
	sp = trimPath(sp)
	t.filePathSlice = sp

	return t.filePathSlice
}
