package methods

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Method struct {
	Name  string
	Class string
}

func ListProjectMethods() []Method {
	content, err := os.ReadFile(path.Join("linter-dot", "functions.html"))
	if err != nil {
		return nil
	}

	reader, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
	if err != nil {
		return nil
	}

	var methods []Method

	selection := reader.Find("div#doc-content > div.contents > ul")
	selection.Children().Map(func(i int, selection *goquery.Selection) string {
		elements := strings.Split(selection.Text(), ":")

		if len(elements) < 2 {
			fmt.Printf("failed to split %v", selection.Text())
			return ""
		}

		methods = append(methods, Method{
			Name:  strings.TrimSpace(elements[0]),
			Class: strings.TrimSpace(elements[1]),
		})
		return ""
	})

	return methods
}
