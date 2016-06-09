package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"sourcegraph.com/sourcegraph/srclib/graph"
	"sourcegraph.com/sourcegraph/srclib/unit"
)

func init() {
	_, err := flagParser.AddCommand("graph",
		"graph man pages",
		"Graph man pages, producing defs of standard commands for srclib-bash.",
		&graphCmd,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Check that we have the '-i' flag.
	cmd := exec.Command("go", "help", "build")
	o, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	usage := strings.Split(string(o), "\n")[0] // The usage is on the first line.
	matched, err := regexp.MatchString("-i", usage)
	if err != nil {
		log.Fatal(err)
	}
	if !matched {
		log.Fatal("'go build' does not have the '-i' flag. Please upgrade to go1.3+.")
	}
}

type GraphCmd struct{}

var graphCmd GraphCmd

func (c *GraphCmd) Execute(args []string) error {
	inputBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("Failed to read STDIN: %s", err)
	}
	var units unit.SourceUnits
	if err := json.NewDecoder(bytes.NewReader(inputBytes)).Decode(&units); err != nil {
		// Legacy API: try parsing input as a single source unit
		var u *unit.SourceUnit
		if err := json.NewDecoder(bytes.NewReader(inputBytes)).Decode(&u); err != nil {
			return fmt.Errorf("Failed to parse source units from input: %s", err)
		}
		units = unit.SourceUnits{u}
	}
	if err := os.Stdin.Close(); err != nil {
		return fmt.Errorf("Failed to close STDIN: %s", err)
	}

	if len(units) == 0 {
		log.Fatal("Input contains no source unit data.")
	}

	out, err := graphUnits(units)
	if err != nil {
		return fmt.Errorf("Failed to graph source units: %s", err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(out); err != nil {
		return fmt.Errorf("Failed to output graph data: %s", err)
	}
	return nil
}

func graphUnits(units unit.SourceUnits) (*graph.Output, error) {
	output := graph.Output{}

	for _, u := range units {
		for _, f := range u.Files {
			graphPage(f, &output)
		}
	}

	return &output, nil
}

func graphPage(page string, output *graph.Output) error {
	f, err := os.Open(page)
	if err != nil {
		return fmt.Errorf("Failed to open file %s: %s", page, err)
	}
	defer f.Close()

	_, file := filepath.Split(page)
	dotIndex := strings.IndexRune(file, '.')
	name := file[:dotIndex]

	def, err := makeCommandDef(page, name, len(name))
	if err != nil {
		return fmt.Errorf("failed to create command def: %s", err)
	}
	output.Defs = append(output.Defs, def)

	return nil
}

func makeCommandDef(filename string, command string, offset int) (*graph.Def, error) {
	data, err := json.Marshal(DefData{
		Name:    command,
		Kind:    "command",
		Keyword: "command",
	})
	if err != nil {
		return nil, err
	}
	return &graph.Def{
		DefKey: graph.DefKey{
			UnitType: "ManPages",
			Unit:     "man",
			Path:     filename + "/" + command,
		},
		Exported: true,
		Data:     data,
		Name:     command,
		Kind:     "command",
		File:     filename,
		DefStart: uint32(offset - len(command)),
		DefEnd:   uint32(offset),
	}, nil
}

type DefData struct {
	Name      string
	Keyword   string
	Type      string
	Kind      string
	Separator string
}
