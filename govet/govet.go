package govet

import (
	"os/exec"
	"strings"

	"github.com/surullabs/lint/checkers"
)

// Check implements a lint.Checker for the govet command.
type Check struct {
	Args []string
}

// Shadow is a Checker that runs
// 	 go tool vet --all --shadow.
var Shadow = Check{Args: []string{"--all", "--shadow"}}

// Check runs go tool vet for pkgs.
func (c Check) Check(pkgs ...string) error {
	files, err := checkers.GoFiles(pkgs...)
	if err != nil {
		return err
	}
	args := append([]string{"tool", "vet"}, append(c.Args, files...)...)
	res, err := checkers.Exec(exec.Command("go", args...))
	if err == nil {
		return nil
	}
	switch res.Code {
	case 1:
		return checkers.Error(strings.Split(strings.TrimSpace(res.Stderr), "\n")...)
	default:
		return err
	}
}
