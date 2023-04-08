package doxygen

import (
	"os"
	"os/exec"
)

func Run() error {
	cmd := exec.Command("doxygen", "Doxyfile")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
