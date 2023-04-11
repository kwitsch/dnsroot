package fileutils

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"regexp"

	"github.com/kwitsch/dnsroot/internal/consts"
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

	if !isCurrentProgramVersion(file.Comments) {
		return "", false
	}

	return getInternicVersion(file.Comments)
}

func isCurrentProgramVersion(comentGroups []*ast.CommentGroup) bool {
	ivr := regexp.MustCompile(`^// dnsroot version: (\d\.\d\.\d)$`)
	for _, cg := range comentGroups {
		for _, c := range cg.List {
			vs := ivr.FindStringSubmatch(c.Text)
			if len(vs) > 1 {
				return (vs[1] == consts.ProgramVersion)
			}
		}
	}
	return false
}

func getInternicVersion(comentGroups []*ast.CommentGroup) (string, bool) {
	ivr := regexp.MustCompile(`^// InterNIC version: (\d+)$`)
	for _, cg := range comentGroups {
		for _, c := range cg.List {
			vs := ivr.FindStringSubmatch(c.Text)
			if len(vs) > 1 {
				return vs[1], true
			}
		}
	}
	return "", false
}
