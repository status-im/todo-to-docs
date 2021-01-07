package main

import (
	"fmt"
	"strings"
)

type todo struct {
	Filepath          string
	Description       string
	LineNumber        int
	RelatedFuncOrType string
	filePathSlice     []string
}

func (t *todo) Path() []string {
	if t.filePathSlice != nil {
		return t.filePathSlice
	}

	sp := strings.Split(t.Filepath, "/")
	sp = t.trimPath(sp)
	t.filePathSlice = sp

	return t.filePathSlice
}

func (t *todo) trimPath(path []string) []string {
	ignoreList := []string{"..", "status-go"}
	if path[0] == ignoreList[0] && path[1] == ignoreList[1] {
		path = path[2:]
	}

	return path
}

func (t *todo) String() string {
	out := ""

	l1 := "*Related func or type* :"
	l2 := "*On line*              :"
	l3 := "*Description*          : "

	if t.RelatedFuncOrType != "" {
		out += l1 + " `" + t.RelatedFuncOrType + "`" + "\n"
	}

	return out +
		fmt.Sprintf("%s %d \n", l2, t.LineNumber) +
		l3 + t.Description
}
