package tmret

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path"
)

var rootPath = "/Users/sawadaikuma/PrivateDev/BackEnd/too-many-results"

func Exec() {
	// ファイルごとのトークンの位置を記録するFileSetを作成する
	fset := token.NewFileSet()

	// ファイル単位で構文解析を行う
	f, err := parser.ParseFile(fset, path.Join(rootPath, "/testdata/src/a/a.go"), nil, 0)
	if err != nil {
		log.Fatal("Error:", err)
	}
	// 抽象構文木を深さ優先で探索する
	ast.Inspect(f, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		// fmt.Printf("pos: %s, type: %s\n", fset.Position(n.Pos()), reflect.TypeOf(n))

		switch a := n.(type) {
		case *ast.FuncDecl:
			rets := a.Type.Results
			if rets == nil {
				return true
			}

			retsExcludeErr := len(rets.List)
			for i := range rets.List {
				ident, ok := rets.List[i].Type.(*ast.Ident)
				if !ok {
					return true
				}

				if ident.Name == "error" {
					retsExcludeErr -= 1
				}

			}
			if retsExcludeErr >= 2 {
				fmt.Println(fset.Position(a.Pos()))
				return false
			}

			return true
		}

		return true
	})
}
