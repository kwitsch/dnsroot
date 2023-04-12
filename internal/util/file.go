package util

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"regexp"
)

const (
	outputFilename = "./dnsroot.go"
)

func OutputFileExists() bool {
	if _, err := os.Stat(outputFilename); err != nil {
		return false
	}

	return true
}

func RemoveOutputFile() error {
	if OutputFileExists() {
		return os.Remove(outputFilename)
	}

	return nil
}

func WriteOutputFile(data []byte) error {
	f, err := os.Create(outputFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func GetCurrentOutputFileVersion() (string, bool) {
	if !OutputFileExists() {
		return "", false
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, outputFilename, nil, parser.ParseComments)
	if err != nil {
		return "", false
	}

	if !isCurrentProgramVersion(file.Comments) {
		return "", false
	}

	return getInternicVersion(file.Comments)
}

func isCurrentProgramVersion(comentGroups []*ast.CommentGroup) bool {
	ivr := regexp.MustCompile(`^// dnsroot version:\s+(\d\.\d\.\d)$`)
	for _, cg := range comentGroups {
		for _, c := range cg.List {
			vs := ivr.FindStringSubmatch(c.Text)
			if len(vs) > 1 {
				return (vs[1] == ProgramVersion)
			}
		}
	}
	return false
}

func getInternicVersion(comentGroups []*ast.CommentGroup) (string, bool) {
	ivr := regexp.MustCompile(`^// InterNIC version:\s+(\d+)$`)
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
