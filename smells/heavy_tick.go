package smells

import (
	"fmt"
	"strings"

	"github.com/quentinguidee/ue-linter/tools"
)

const (
	StateSearchingFor = iota
	StateSearchingOpeningBrace
	StateSearchingClosingBrace
)

type HeavyTickAnalyzer struct {
	Methods    map[string]tools.Method
	CallGraphs *[]tools.CallGraph
}

func (a HeavyTickAnalyzer) Run() error {
	var ticks []tools.Method

	for _, graph := range *a.CallGraphs {
		if !strings.Contains(graph.Name, "Tick") {
			continue
		}

		tick := a.Methods[graph.Name[1:len(graph.Name)-1]]

		var children []tools.Method

		for _, node := range (*graph.Nodes).Nodes {
			id := node.Attrs["label"]
			id = id[1 : len(id)-1]
			if !strings.Contains(id, "Tick") {
				children = append(children, a.Methods[id])
			}
		}

		changed := true
		maxDepth := 100
		depth := 0

		for changed && depth < maxDepth {
			changed = false

			//println("---")
			//println(tick.Content)

			for _, m := range children {
				if strings.Count(tick.Content, m.Name+"(") > 0 {
					tick.Content = strings.ReplaceAll(tick.Content, m.Name, m.Content)
					changed = true
				}
			}

			depth += 1
		}

		if depth >= maxDepth {
			fmt.Printf("WARN: Max recursion depth of %d exceeded for the heavy_tick analyzer.\n", maxDepth)
		}

		ticks = append(ticks, tick)
	}

	for _, tick := range ticks {
		a.FindInMethod(tick)
	}

	return nil
}

func (a HeavyTickAnalyzer) FindInMethod(m tools.Method) {
	level := 0
	maxLevel := 0

	content := m.Content
	i := 0

	state := StateSearchingFor

	for len(content) > i {
		switch state {
		case StateSearchingFor:
			if len(content) > i+2 && content[i:i+3] == "for" {
				i += 2
				state = StateSearchingOpeningBrace
			}
			break

		case StateSearchingOpeningBrace:
			if content[i] == '{' {
				level += 1
				if maxLevel < level {
					maxLevel = level
				}

				state = StateSearchingClosingBrace
			}
			break

		case StateSearchingClosingBrace:
			if content[i] == '}' {
				level -= 1
			} else if len(content) > i+2 && content[i:i+3] == "for" {
				i += 2
				state = StateSearchingOpeningBrace
			}
			break
		}

		i += 1
	}

	if maxLevel >= 3 {
		fmt.Printf("Potential game smell in %s::%s: found a loop with a depth of %d.\n", m.Class, m.Name, maxLevel)
	}
}
