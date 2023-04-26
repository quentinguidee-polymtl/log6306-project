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

	for _, m := range methodsToAnalyze {
		a.FindInMethod(m)
	}

	return nil
}

func (a TickSpawnDestroyAnalyzer) FindInMethod(m tools.Method) {
	countSpawn := strings.Count(m.Content, "SpawnActor")
	if countSpawn > 0 {
		fmt.Printf("Potential game smell in %s::%s: found %d SpawnActor call.\n", m.Class, m.Name, countSpawn)
	}

	countDestroy := strings.Count(m.Content, "DestroyActor")
	if countDestroy > 0 {
		fmt.Printf("Potential game smell in %s::%s: found %d DestroyActor call.\n", m.Class, m.Name, countDestroy)
	}
}
