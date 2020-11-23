package lint

import (
	"golang.org/x/tools/go/analysis"
	"strconv"
)

var Analyzer = &analysis.Analyzer{
	Name: "import_unsafe",
	Doc:  "reports imports of package unsafe",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, imp := range f.Imports {
			path, err := strconv.Unquote(imp.Path.Value)
			if err == nil && path == "unsafe" {
				pass.Reportf(imp.Pos(), "package unsafe must not be imported")
			}
		}
	}
	return nil, nil
}
