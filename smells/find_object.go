package smells

import (
	"fmt"
	"regexp"

	"github.com/quentinguidee/ue-linter/tools"
)

type FindObjectAnalyzer struct {
	Methods map[string]tools.Method
}

func (a FindObjectAnalyzer) Run() error {
	r := regexp.MustCompile(`FindObject[(<]`)

	for _, m := range a.Methods {
		matches := r.FindStringSubmatch(m.Content)

		if len(matches) > 0 {
			fmt.Printf("Potential game smell in %s::%s: found %d FindObject call.\n", m.Class, m.Name, len(matches))
		}
	}

	return nil
}
