package dot

import (
	"github.com/awalterschulze/gographviz"
	"os"
	"path"
	"strings"
)

func ParseAll() ([]*gographviz.Graph, error) {
	var graphs []*gographviz.Graph

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

			graphs = append(graphs, graph)
		}
	}

	return graphs, nil
}
