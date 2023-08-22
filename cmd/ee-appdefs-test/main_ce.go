//go:build !ee

package main

import "fmt"

func run() error {
	fmt.Println("Nothing to see here")
	return nil
}
