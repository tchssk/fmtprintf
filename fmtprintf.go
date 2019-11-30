package fmtprintf

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/types/typeutil"
)

const Doc = `report fmt.Printf call which have one argument`

var Analyzer = &analysis.Analyzer{
	Name:     "fmtprintf",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		call := n.(*ast.CallExpr)

		fn := typeutil.StaticCallee(pass.TypesInfo, call)
		if fn == nil {
			return
		}
		if fn.Name() != "Printf" {
			return
		}
		if fn.Type().(*types.Signature).Recv() != nil {
			return
		}
		if fn.Pkg().Path() != "fmt" {
			return
		}
		if len(call.Args) > 1 {
			return
		}

		selector, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}

		pass.Report(analysis.Diagnostic{
			Pos:     call.Pos(),
			Message: "fmt.Printf call which have one argument can be replaced with fmt.Print",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "Replace",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     selector.Sel.Pos(),
							End:     selector.Sel.End(),
							NewText: []byte("Print"),
						},
					},
				},
			},
		})
	})

	return nil, nil
}
