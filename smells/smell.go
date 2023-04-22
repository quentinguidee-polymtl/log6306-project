package smells

import (
	"io/fs"
	"path"
	"path/filepath"

	"github.com/awalterschulze/gographviz"
)

type SmellAnalyzer interface {
	Run(graphs []*gographviz.Graph) error
}

func walk(run func(filepath string)) error {
	err := filepath.Walk(".", func(filepath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path.Ext(info.Name()) == ".cpp" || path.Ext(info.Name()) == ".h" {
			run(filepath)
		}

		return err
	})
	return err
}
