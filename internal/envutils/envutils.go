package envutils

import "os"

const (
	gofile   = "GOFILE"
	gomodule = "GOPACKAGE"
)

func GetGoFile() (string, bool) {
	return os.LookupEnv(gofile)
}

func GetGoModule() (string, bool) {
	return os.LookupEnv(gomodule)
}
