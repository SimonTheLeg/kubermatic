/*
Copyright 2023 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Config struct {
	GoModDir   string
	Provider   string
	Pkgs       []string
	Filter     *regexp.Regexp
	PrintFuncs bool

	PoC PolicyCreator
}

var defaultConfig = Config{
	Provider:   "", // must be set by the user
	GoModDir:   "", // will be defaulted to working directory in Init
	PrintFuncs: false,
}

var supportedProviders = []string{"aws"}

type InvalidConfigError struct{}

func (e InvalidConfigError) Error() string { return "Invalid Config: " }

func (c *Config) InitAndValidate() error {
	if c.Provider == "" {
		return fmt.Errorf("Flag --provider must be set. Supported providers are: %s", supportedProviders)
	}

	if c.Provider != "aws" {
		return fmt.Errorf("Provider '%s' is not supported. Supported providers are: %s", c.Provider, supportedProviders)
	}

	if c.GoModDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("Error determining current working directory: %w", err)
		}
		c.GoModDir = wd
	}

	pkgs := flag.Arg(0)
	if pkgs == "" {
		return fmt.Errorf("You need to provide packages and a filter to search for")
	}
	c.Pkgs = strings.Split(pkgs, ",")

	filter := flag.Arg(1)
	if filter == "" {
		return fmt.Errorf("You need to provide packages and a filter to search for")
	}
	regex, err := regexp.Compile(filter)
	if err != nil {
		return fmt.Errorf("Provided filter '%s', is not a valid regex: %w", filter, err)
	}
	c.Filter = regex

	if c.Provider == "aws" {
		mapper, err := NewAWSDefaultMapper()
		if err != nil {
			return err
		}
		c.PoC = NewAWSPolicyCreator(mapper)
	}

	return nil
}

func (c *Config) FromFlags() {
	flag.StringVar(&c.Provider, "provider", defaultConfig.Provider, fmt.Sprintf("Cloud provider to generate policy for. Must be set. Supported providers: %s", supportedProviders))
	flag.StringVar(&c.GoModDir, "modDir", defaultConfig.GoModDir, "The base path from which to start searching. This should be where your go.mod file is located. Defaults to current directory")
	flag.BoolVar(&c.PrintFuncs, "printFuncs", defaultConfig.PrintFuncs, "Whether you only want to print the func usage, and not generate a whole provider statement. Defaults to false")
}
