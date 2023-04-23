package smells

import (
	"log"

	methodstool "github.com/quentinguidee/ue-linter/tools/methods"
)

type TickSpawnDestroyAnalyzer struct {
	Methods []*methodstool.Method
}

func (t TickSpawnDestroyAnalyzer) Run() error {
	//var ticks []methodstool.Method
	//
	//for _, method := range t.Methods {
	//	if method.Name == "Tick()" {
	//		ticks = append(ticks, *method)
	//	}
	//}

	tickMethods, err := getTickMethods()
	if err != nil {
		return err
	}

	log.Printf("%+v", tickMethods)

	return err
}
