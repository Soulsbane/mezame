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

func (d *DrivePaths) Add(path string) {
	d.paths.Put(path, true)
}

func (d *DrivePaths) Remove(path string) {
	d.paths.Remove(path)
}

func (d *DrivePaths) Enable(path string) {
	d.paths.Put(path, true)
}

func (d *DrivePaths) Disable(path string) {
	d.paths.Put(path, false)
}

func (d *DrivePaths) NumPaths() int {
	return d.paths.Size()
}

func (d *DrivePaths) Dump() {
	for _, key := range d.paths.Keys() {
		value, _ := d.paths.Get(key)
		fmt.Println(key, value)
	}
}
