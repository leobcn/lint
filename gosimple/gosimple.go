package gosimple

import "github.com/surullabs/statictest/checkers"

// Check implements a gosimple Checker (https://github.com/dominikh/go-simple)
type Check struct {
}

// Check runs gosimple for pkg
func (Check) Check(pkgs ...string) error {
	return checkers.LintAsFiles("gosimple", "honnef.co/go/simple/cmd/gosimple", pkgs)
}