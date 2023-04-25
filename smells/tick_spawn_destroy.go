package smells

import (
	"fmt"
	"strings"

	"github.com/quentinguidee/ue-linter/tools"
)

type TickSpawnDestroyAnalyzer struct {
	Methods    map[string]tools.Method
	CallGraphs *[]tools.CallGraph
}

func (a TickSpawnDestroyAnalyzer) Run() error {
	// Find all methods executed during ticks
	var methodsToAnalyze []tools.Method

	for _, graph := range *a.CallGraphs {
		if !strings.Contains(graph.Name, "Tick") {
			continue
		}
		for _, node := range (*graph.Nodes).Nodes {
			id := node.Attrs["label"]
			id = id[1 : len(id)-1]
			methodsToAnalyze = append(methodsToAnalyze, a.Methods[id])
		}
	}

	for i, m := range methodsToAnalyze {
		fmt.Printf("=> ANALYZER_TICK_SPAWN_DESTROY %d: %s::%s\n", i+1, m.Class, m.Name)
		a.FindInMethod(m)
	}

	return nil
}

func (a TickSpawnDestroyAnalyzer) FindInMethod(m tools.Method) {
	if strings.Contains(m.Content, "SpawnActor") {
		fmt.Printf("Potential game smell in %s::%s: found a SpawnActor call.\n", m.Class, m.Name)
	}

	if strings.Contains(m.Content, "DestroyActor") {
		fmt.Printf("Potential game smell in %s::%s: found a DestroyActor call.\n", m.Class, m.Name)
	}
}
