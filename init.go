package main

import (
	"flag"
	"log"
	"os"

	"github.com/awalterschulze/gographviz"
	"github.com/quentinguidee/ue-linter/smells"
	"github.com/quentinguidee/ue-linter/tools/dot"
	"github.com/quentinguidee/ue-linter/tools/doxygen"
	methodstool "github.com/quentinguidee/ue-linter/tools/methods"
)

func main() {
	parseArgs()

	if false {
		err := doxygen.Run()
		if err != nil {
			os.Exit(1)
		}
	}

	graphs, err := dot.ParseAll()
	if err != nil {
		log.Fatalln(err)
	}

	methods := methodstool.ListProjectMethods()
	log.Printf("%v", methods)

	err = analyze(graphs)
	if err != nil {
		log.Fatalln(err)
	}
}

func analyze(graphs []*gographviz.Graph) error {
	smellAnalyzers := []smells.SmellAnalyzer{
		smells.TickSpawnDestroyAnalyzer{},
	}

	for _, smellAnalyzer := range smellAnalyzers {
		err := smellAnalyzer.Run(graphs)
		if err != nil {
			return err
		}
	}

	return nil
}

func parseArgs() {
	projectPath := flag.String("path", ".", "Define the UE project path.")
	flag.Parse()
	err := os.Chdir(*projectPath)
	if err != nil {
		log.Fatalf("failed to change the workdir: %v", err)
	}
}
