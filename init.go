package main

import (
	"flag"
	"log"
	"os"

	"github.com/quentinguidee/ue-linter/smells"
	"github.com/quentinguidee/ue-linter/tools"
)

func main() {
	parseArgs()

	if false {
		err := tools.RunDoxygen()
		if err != nil {
			os.Exit(1)
		}
	}

	graphs, err := tools.ParseAll()
	if err != nil {
		log.Fatalln(err)
	}

	methods, err := tools.ListProjectMethods()
	if err != nil {
		log.Fatalln(err)
	}

	err = analyze(methods, &graphs)
	if err != nil {
		log.Fatalln(err)
	}
}

func analyze(methods map[string]tools.Method, callGraphs *[]tools.CallGraph) error {
	smellAnalyzers := []smells.SmellAnalyzer{
		smells.TickSpawnDestroyAnalyzer{
			Methods:    methods,
			CallGraphs: callGraphs,
		},
		smells.FindObjectAnalyzer{
			Methods: methods,
		},
		smells.HeavyTickAnalyzer{
			Methods:    methods,
			CallGraphs: callGraphs,
		},
	}

	for _, smellAnalyzer := range smellAnalyzers {
		err := smellAnalyzer.Run()
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
