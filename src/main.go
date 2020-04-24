package main

import (
	"awesomeGvm/src/classpath"
	"fmt"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Print("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		usage()
	} else {
		startJvm(cmd)
	}
}

func startJvm(cmd *Cmd) {
	cp := classpath.Parse(cmd.xJreOption, cmd.cpOption)
	fmt.Printf("classpath: %v class %v args:%v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
