package lion

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/analysisutil"
	"github.com/gostaticanalysis/coverprofile"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "lion",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		coverprofile.Analyzer,
	},
}

const Doc = "lion finds functions which are not tested"

func run(pass *analysis.Pass) (interface{}, error) {
	// ignore test files
	if pkg := pass.Pkg.Path(); strings.HasSuffix(pkg, ".test") ||
		strings.HasSuffix(pkg, "_test") {
		return nil, nil
	}

	files := pass.ResultOf[coverprofile.Analyzer].([]*coverprofile.File)

	for _, f := range getAllFuncs(pass) {
		if !isTested(pass, f, files) {
			pass.Reportf(f.Pos(), "%s is not tested", f.FullName())
		}
	}

	return nil, nil
}

func getAllFuncs(pass *analysis.Pass) []*types.Func {
	var fs []*types.Func

LOOP:
	for _, n := range pass.Pkg.Scope().Names() {
		obj := pass.Pkg.Scope().Lookup(n)
		switch obj := obj.(type) {
		case *types.Func:
			fs = append(fs, obj)
		case *types.TypeName:
			typ, ok := obj.Type().(*types.Named)
			if !ok {
				break LOOP
			}

			for i := 0; i < typ.NumMethods(); i++ {
				fs = append(fs, typ.Method(i))
			}
		}
	}

	return fs
}

func isTested(pass *analysis.Pass, f *types.Func, files []*coverprofile.File) bool {
	decl := getFuncDecl(pass, f.Pos())

	for _, file := range files {
		if file.TokenFile != tokenFile(pass, f.Pos()) {
			continue
		}

		for _, b := range file.Blocks {
			if b.Count == 0 {
				continue
			}
			if (decl.Pos() <= b.Start && b.Start <= decl.End()) ||
				(decl.Pos() <= b.End && b.End <= decl.End()) {
				return true
			}
		}
	}

	return false
}

func tokenFile(pass *analysis.Pass, pos token.Pos) *token.File {
	var file *token.File
	p := int(pos)
	pass.Fset.Iterate(func(f *token.File) bool {
		if f.Base() <= p && p <= f.Base()+f.Size() {
			file = f
			return false
		}
		return true
	})
	return file
}

func getFuncDecl(pass *analysis.Pass, pos token.Pos) *ast.FuncDecl {
	f := analysisutil.File(pass, pos)
	if f == nil {
		return nil
	}

	var funcdecl *ast.FuncDecl
	ast.Inspect(f, func(n ast.Node) bool {
		d, ok := n.(*ast.FuncDecl)
		if ok && d.Pos() <= pos && pos <= d.End() {
			funcdecl = d
			return false
		}
		return true
	})

	return funcdecl
}
