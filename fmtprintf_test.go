package fmtprintf_test

import (
	"testing"

	"github.com/tchssk/fmtprintf"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, fmtprintf.Analyzer, "a")
}
