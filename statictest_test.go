package statictest_test

import (
	"testing"

	"github.com/surullabs/statictest"
	"github.com/surullabs/statictest/gofmt"
	"github.com/surullabs/statictest/golint"
	"github.com/surullabs/statictest/govet"
)

func TestStaticChecks(t *testing.T) {
	if err := statictest.Chain(gofmt.Check{}, govet.Shadow, golint.Check{}).Check("."); err != nil {
		t.Fatal(err)
	}
}
