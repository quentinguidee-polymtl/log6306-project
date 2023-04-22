package smells

import (
	"os"

	"github.com/awalterschulze/gographviz"
)

type TickSpawnDestroyAnalyzer struct{}

func (t TickSpawnDestroyAnalyzer) Run(graphs []*gographviz.Graph) error {
	err := walk(func(filepath string) {
		_, err := os.ReadFile(filepath)
		if err != nil {
			return
		}

		//code := string(content)

		//println(code)
	})
	return err
}
