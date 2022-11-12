package drivepaths

import (
	"fmt"

	"github.com/daichi-m/go18ds/maps/hashmap"
)

type DrivePaths struct {
	pathsFileName string
	paths         *hashmap.Map[string, bool]
}

func New(fileName string) *DrivePaths {
	var locs DrivePaths

	locs.paths = hashmap.New[string, bool]()
	locs.pathsFileName = fileName

	return &locs
}

func (f *DrivePaths) AddPath(path string) {
	f.paths.Put(path, true)
}

func (f *DrivePaths) RemovePath(path string) {
	f.paths.Remove(path)
}

func (f *DrivePaths) EnablePath(path string) {
	f.paths.Put(path, true)
}

func (f *DrivePaths) DisablePath(path string) {
	f.paths.Put(path, false)
}

func (f *DrivePaths) Dump() {
	for _, key := range f.paths.Keys() {
		value, _ := f.paths.Get(key)
		fmt.Println(key, value)
	}
}
