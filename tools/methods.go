package tools

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

type Method struct {
	Name    string
	Class   string
	Content string
}

func walk(run func(filepath string)) error {
	err := filepath.Walk(".", func(filepath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path.Ext(info.Name()) == ".cpp" /* || path.Ext(info.Name()) == ".h" */ {
			run(filepath)
		}
		return err
	})
	return err
}

func ListProjectMethods() (map[string]Method, error) {
	var methods = map[string]Method{}

	err := walk(func(filepath string) {
		file, err := os.Open(filepath)
		if err != nil {
			return
		}

		fileScanner := bufio.NewScanner(file)

		match := 0
		inTick := false
		method := Method{}
		for fileScanner.Scan() {
			line := fileScanner.Text()

			if !inTick {
				r := regexp.MustCompile(`(.*)::(.*)\((.*)\)`)
				matches := r.FindStringSubmatch(line)
				if len(matches) != 0 {
					// Start tick
					inTick = true
					elements := strings.Split(matches[0], "::")
					left := strings.Split(elements[0], " ")
					right := strings.Split(elements[1], " ")
					right = strings.Split(right[0], "(")
					method.Class = left[len(left)-1]
					method.Name = right[0]
				}
			}

			if inTick {
				openBracketsCount := strings.Count(line, "{")
				closeBracketsCount := strings.Count(line, "}")

				match += openBracketsCount - closeBracketsCount

				line = strings.TrimSpace(line)
				if line != "" && !strings.HasPrefix(line, "//") {
					method.Content += line + "\n"
				}

				if match <= 0 && openBracketsCount+closeBracketsCount > 0 {
					// End tick
					inTick = false
					match = 0

					// Keep only the method body
					method.Content = method.Content[strings.Index(method.Content, "{")+1 : strings.LastIndex(method.Content, "}")]
					method.Content = strings.TrimSpace(method.Content)

					methods[fmt.Sprintf(method.Class+"::"+method.Name)] = method

					method = Method{}
				}
			}
		}
	})

	return methods, err
}
