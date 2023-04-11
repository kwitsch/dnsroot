package util

import "os"

const (
	gopackage = "GOPACKAGE"
)

func GetGoPackage() (string, bool) {
	return os.LookupEnv(gopackage)
}
