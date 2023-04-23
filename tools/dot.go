package tools

import (
	"os"
	"path"
	"strings"

	"github.com/awalterschulze/gographviz"
)

type CallGraph struct {
	gographviz.Graph
}

func ParseAll() ([]CallGraph, error) {
	var graphs []CallGraph

	entries, err := os.ReadDir("linter-dot")
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		p := path.Join("linter-dot", entry.Name())
		if strings.HasSuffix(entry.Name(), ".dot") {
			file, err := os.ReadFile(p)
			if err != nil {
				return nil, err
			}

			graphAst, err := gographviz.Parse(file)
			if err != nil {
				return nil, err
			}

			graph := gographviz.NewGraph()
			err = gographviz.Analyse(graphAst, graph)
			if err != nil {
				return nil, err
			}

			graphs = append(graphs, CallGraph{*graph})
		}
	}

	return graphs, nil
}

func (g CallGraph) Get(method string) error {
	return nil
}
