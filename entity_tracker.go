package main

import (
	"regexp"
)

type entityTracker struct {
	entityDetector       *regexp.Regexp
	openBracketDetector  *regexp.Regexp
	closeBracketDetector *regexp.Regexp

	currentEntity    string
	lastEntity       string
	openBracketCount map[string]int
}

func NewEntityTracker() (*entityTracker, error) {
	ed, err := regexp.Compile("(func.+\\(.*\\).*|type .+) \\{")
	if err != nil {
		return nil, err
	}

	obd, err := regexp.Compile("\\{")
	if err != nil {
		return nil, err
	}

	cbd, err := regexp.Compile("\\}")
	if err != nil {
		return nil, err
	}

	return &entityTracker{
		entityDetector:       ed,
		openBracketDetector:  obd,
		closeBracketDetector: cbd,
		openBracketCount:     map[string]int{},
	}, nil
}

func (et *entityTracker) Track(line string) {
	ft := et.entityDetector.FindStringSubmatch(line)
	if ft != nil {
		// If a struct is declared inside a function
		if et.openBracketCount[et.currentEntity] > 0 {
			et.lastEntity = et.currentEntity
		}

		et.currentEntity = ft[1]
		et.openBracketCount[et.currentEntity]++
		return
	}

	ob := et.openBracketDetector.FindAllString(line, -1)
	if ob != nil {
		et.openBracketCount[et.currentEntity] += len(ob)
	}

	cb := et.closeBracketDetector.FindAllString(line, -1)
	if cb != nil {
		et.openBracketCount[et.currentEntity] -= len(cb)

		// If current entity has closed its final bracket delete the tracking from map
		// then check if there is an active lastEntity, if so replace currentEntity with lastEntity
		if et.openBracketCount[et.currentEntity] == 0 {
			delete(et.openBracketCount, et.currentEntity)

			if et.lastEntity != "" {
				et.currentEntity = et.lastEntity
				et.lastEntity = ""
			}
		}
	}

}

func (et entityTracker) Current() string {
	if et.openBracketCount[et.currentEntity] == 0 {
		return ""
	}

	return et.currentEntity
}
