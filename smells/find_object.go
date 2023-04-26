package smells

import (
	"fmt"
	"regexp"

	"github.com/quentinguidee/ue-linter/tools"
)

var r = regexp.MustCompile(`FindObject[(<]`)

type FindObjectAnalyzer struct {
	Methods map[string]tools.Method
}

func (a FindObjectAnalyzer) Run() error {
	for _, m := range a.Methods {
		a.FindInMethod(m)
	}
	return nil
}

func (a FindObjectAnalyzer) FindInMethod(m tools.Method) {
	matches := r.FindStringSubmatch(m.Content)
	if len(matches) > 0 {
		fmt.Printf("Potential game smell in %s::%s: found %d FindObject call.\n", m.Class, m.Name, len(matches))
	}
}
