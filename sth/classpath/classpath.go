package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClassPath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (c *Classpath) ReadClass(classname string) ([]byte, Entry, error) {
	classname = classname + ".class"
	if class, entry, err := c.bootClassPath.ReadClass(classname); err == nil {
		return class, entry, err
	}
	if class, entry, err := c.extClasspath.ReadClass(classname); err == nil {
		return class, entry, err
	}
	return c.userClasspath.ReadClass(classname)
}

func (c *Classpath) String() string {
	return c.userClasspath.String()
}

func (c *Classpath) parseBootAndExtClasspath(option string) {
	jreDir := getJreDir(option)

	//jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	c.bootClassPath = newWildcardEntry(jreLibPath)

	// jre.lib.ext.*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	c.extClasspath = newWildcardEntry(jreExtPath)
}

func (c *Classpath) parseUserClasspath(option string) {
	if option == "" {
		option = "."
	}
	c.userClasspath = newEntry(option)
}

func getJreDir(option string) string {
	if option != "" && exists(option) {
		return option
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("can not find jre folder!")
}

func exists(option string) bool {
	if _, err := os.Stat(option); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
