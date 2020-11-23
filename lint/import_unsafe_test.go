package lint

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestAll(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer)
}
