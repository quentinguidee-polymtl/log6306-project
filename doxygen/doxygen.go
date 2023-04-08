package doxygen

import (
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
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
	entries, err := os.ReadDir("linter-dot")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		ext := filepath.Ext(entry.Name())

		if ext != ".dot" {
			err := os.Remove(path.Join("linter-dot", entry.Name()))
			if err != nil {
				log.Fatalf("%v", err)
				return err
			}
		}
	}

	return nil
}
