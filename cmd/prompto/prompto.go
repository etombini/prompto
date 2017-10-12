package main

import (
	"flag"
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

	config := flag.String("config", "~/.prompto.yaml", "Configuration file")
	help := flag.Bool("help", false, "Get some help")
	shell := flag.String("shell", "bash", "Define the shell used (bash, zsh, fish)")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	prompter.SetShell(*shell)
	lines := loadYAML(*config)

	for _, line := range lines.Lines {
		//l := line.String()
		fmt.Printf("%s", line)
	}
}
