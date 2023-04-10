package fileutils

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
)

const (
	filename = "./dnsroot.go"
)

func GetCurrentVersion() (string, bool) {
	if _, err := os.Stat(filename); err != nil {
		return "", false
	}

	fset := token.NewFileSet()
	//Parse the file and create an AST
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return "", false
	}

	for _, cg := range file.Comments {
		for _, c := range cg.List {
			fmt.Println(c.Text)
		}
	}
	return "", true
}
