//go:generate go run .
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kwitsch/dnsroot/internal/rootfile"
	"github.com/kwitsch/dnsroot/internal/util"
)

const (
	gofile = "GOFILE"
)

func main() {
	log.SetOutput(os.Stdout)
	_, isSet := util.GetGoPackage()
	if !isSet {
		fmt.Println("GOPACKAGE is missing in environment")
		fmt.Println("This application is intended to be used with go:generate")
		os.Exit(1)
	}

	vs, found := util.GetCurrentOutputFileVersion()
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

	if vs != cf.Version {
		fmt.Println("has to be updated")
	}
}
