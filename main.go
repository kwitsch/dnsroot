//go:generate go run .
package main

import (
	"fmt"
	"os"

	"github.com/kwitsch/dnsroot/internal/generator"
	"github.com/kwitsch/dnsroot/internal/util"
)

func main() {
	if _, isSet := util.GetGoPackage(); !isSet {
		fmt.Println("GOPACKAGE is missing in environment")
		fmt.Println("This application is intended to be used with go:generate")
		os.Exit(1)
	}

	if err := generator.Run(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
