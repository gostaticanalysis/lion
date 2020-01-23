package lion_test

import (
	"testing"

	"github.com/gostaticanalysis/lion"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, lion.Analyzer, "a")
}