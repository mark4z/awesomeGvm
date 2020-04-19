package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) ReadClass(classname string) ([]byte, Entry, error) {
	filename := filepath.Join(self.absDir, classname)
	data, err := ioutil.ReadFile(filename)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
