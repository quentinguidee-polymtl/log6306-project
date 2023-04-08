package main

import (
	"flag"
	"log"
	"os"

	"github.com/quentinguidee/ue-linter/doxygen"
)

func main() {
	parseArgs()

	err := doxygen.Run()
	if err != nil {
		os.Exit(1)
	}
}

func parseArgs() {
	projectPath := flag.String("path", ".", "Define the UE project path.")
	flag.Parse()
	err := os.Chdir(*projectPath)
	if err != nil {
		log.Fatalf("failed to change the workdir: %v", err)
	}
}
