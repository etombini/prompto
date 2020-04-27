package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/etombini/prompto/pkg/prompter"
	yaml "gopkg.in/yaml.v2"
)

//PromptoConfig is the datastructure holding configuration for each component
type PromptoConfig struct {
	Prompters []map[string]string
}

func loadConfig(f string) (PromptoConfig, error) {
	config, err := ioutil.ReadFile(f)
	if err != nil {
		return PromptoConfig{}, err
	}

	c := PromptoConfig{}
	if err := yaml.Unmarshal(config, &c); err != nil {
		return PromptoConfig{}, err
	}

	return c, nil
}

func usage() {
	fmt.Printf("Usage : prompto configuration_file\n")
}

var help = map[string]bool{
	"-h":     true,
	"-help":  true,
	"--help": true,
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	if _, ok := help[os.Args[1]]; ok {
		usage()
		os.Exit(0)
	}

	y, err := loadConfig(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "can not load configuration from %s: %s\n", os.Args[1], err)
		os.Exit(1)
	}
	p := prompter.New(y.Prompters)
	fmt.Printf("%s", p.String())
}
