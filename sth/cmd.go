package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
	xJreOption  string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = usage

	flag.BoolVar(&cmd.versionFlag, "version", false, "the version")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.xJreOption, "jre", "", "path 2 jre")

	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func usage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
