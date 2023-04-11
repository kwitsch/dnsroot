//go:generate go run .
package main

import (
	"fmt"
	"os"

	"github.com/kwitsch/dnsroot/internal/envutils"
	"github.com/kwitsch/dnsroot/internal/fileutils"
	"github.com/kwitsch/dnsroot/internal/rootfile"
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
	vs, found := fileutils.GetCurrentVersion()
	fmt.Println(vs, found)

	cf, err := rootfile.Get()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println("Last Update:", cf.LastUpdate)
	fmt.Println("Version:", cf.Version)
	fmt.Println("Servers:")
	for _, s := range cf.Servers {
		fmt.Println(s)
	}
}
