package main

import (
	"github.com/tchssk/fmtprintf"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(fmtprintf.Analyzer) }
