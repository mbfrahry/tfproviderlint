// Package S007 defines an Analyzer that checks for
// Schema with Required being false
package S007

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/schemaschema"
)

const Doc = `check for Schema with Required being false

The S007 analyzer reports cases of schemas which includes Required
being false, which is an unnecessary  assignment.`

const analyzerName = "S007"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		schemaschema.Analyzer,
		commentignore.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemas := pass.ResultOf[schemaschema.Analyzer].([]*ast.CompositeLit)
	for _, schema := range schemas {
		if ignorer.ShouldIgnore(analyzerName, schema) {
			continue
		}

		var requiredDisabled bool

		for _, elt := range schema.Elts {
			switch v := elt.(type) {
			default:
				continue
			case *ast.KeyValueExpr:
				name := v.Key.(*ast.Ident).Name

				if name != "Required" {
					continue
				}

				switch v := v.Value.(type) {
				default:
					continue
				case *ast.Ident:
					value := v.Name

					if value != "true" {
						continue
					}

					requiredDisabled = true
				}
			}
		}

		if requiredDisabled {
			pass.Reportf(schema.Type.(*ast.SelectorExpr).Sel.Pos(), "%s: schema should not include `Required: false,`", analyzerName)
		}
	}

	return nil, nil
}
