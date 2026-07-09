//go:build ignore

package main

import (
	"fmt"
	"os"

	"github.com/seakee/cpa-manager-plus/apps/manager-server/internal/security"
)

func main() {
	keyPath := os.Args[1]
	encrypted := os.Args[2]

	dataKey, _, err := security.LoadOrCreateDataKey("", keyPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load data key: %v\n", err)
		os.Exit(1)
	}
	p, err := security.NewProtector(dataKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "new protector: %v\n", err)
		os.Exit(1)
	}
	plain, err := p.UnprotectString(encrypted)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unprotect: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(plain)
}
