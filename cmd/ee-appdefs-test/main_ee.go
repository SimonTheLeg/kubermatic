//go:build ee

package main

import (
	"fmt"

	appdefs "github.com/kubermatic-labs/ee-appdefs-test"
)

func run() error {
	files, err := appdefs.GetAppDefFiles()
	if err != nil {
		return err
	}

	for _, file := range files {
		finfo, err := file.Stat()
		if err != nil {
			return err
		}
		fmt.Println(finfo.Name())
	}
	return nil
}
