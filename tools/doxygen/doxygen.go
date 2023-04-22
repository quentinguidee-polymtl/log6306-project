package doxygen

import (
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func Run() error {
	cmd := exec.Command("doxygen", "Doxyfile")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		return err
	}

	err = cleanup()
	return err
}

// cleanup deletes all useless files
func cleanup() error {
	doxygenPath := "linter-dot"

	entries, err := os.ReadDir(doxygenPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		p := path.Join(doxygenPath, entry.Name())

		if entry.IsDir() {
			err = os.RemoveAll(p)
			if err != nil {
				return err
			}
			continue
		}

		if !strings.HasSuffix(entry.Name(), "cgraph.dot") && !strings.HasSuffix(entry.Name(), "cgraph.png") && entry.Name() != "functions.html" {
			err := os.Remove(p)
			if err != nil {
				log.Fatalf("%v", err)
				return err
			}
		}
	}

	return nil
}
