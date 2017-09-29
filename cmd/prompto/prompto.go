package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/etombini/prompto/pkg/prompter"
	yaml "gopkg.in/yaml.v2"
)

func loadYAML(filename string) *prompter.Config {
	config, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can not read %s - %s\n", filename, err.Error())
		os.Exit(1)
	}

	var cfg = prompter.Config{}

	if err := yaml.Unmarshal(config, &cfg); err != nil {
		fmt.Printf("Problem unmarshalling: %s\n", err.Error())
		os.Exit(1)
	}

	return &cfg
}

func main() {
	lines := loadYAML(os.Args[1])

	for _, line := range lines.Lines {
		l := line.String()
		fmt.Printf("%s", l)
	}
}
