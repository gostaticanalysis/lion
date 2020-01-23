package main

import (
	"github.com/gostaticanalysis/lion"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(lion.Analyzer) }
