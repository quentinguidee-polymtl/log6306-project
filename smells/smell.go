package smells

import (
	"bufio"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

type SmellAnalyzer interface {
	Run() error
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

type Tick struct {
	Class   string
	Content string
}

func getTickMethods() ([]Tick, error) {
	var tickBodies []Tick

	err := walk(func(filepath string) {
		file, err := os.Open(filepath)
		if err != nil {
			return
		}

		fileScanner := bufio.NewScanner(file)

		match := 0
		inTick := false
		class := ""
		body := ""
		for fileScanner.Scan() {
			line := fileScanner.Text()

			if !inTick {
				r := regexp.MustCompile(` (.*)::Tick`)
				calls := r.FindStringSubmatch(line)
				if len(calls) != 0 {
					// Start tick
					inTick = true
					class = strings.Split(calls[0], "::")[0]
				}
			}

			if inTick {
				change := 0
				change += strings.Count(line, "{")
				change -= strings.Count(line, "}")

				body += line

				if match > 0 && match+change <= 0 {
					// End tick
					inTick = false
					match = 0
					tickBodies = append(tickBodies, Tick{
						Class:   class,
						Content: body,
					})
					body = ""
				}

				match += change
			}
		}
	})

	return tickBodies, err
}
