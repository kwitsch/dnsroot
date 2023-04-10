//go:generate go run .
package main

import (
	"fmt"
	"os"

	"github.com/kwitsch/dnsroot/internal/envutils"
	"github.com/kwitsch/dnsroot/internal/fileutils"
)

const (
	gofile = "GOFILE"
)

func main() {
	envVar, isSet := envutils.GetGoModule()
	if !isSet {
		os.Exit(1)
	}
	fmt.Println("Environment =", envVar)
	fileutils.GetCurrentVersion()
}
