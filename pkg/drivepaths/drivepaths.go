package drivepaths

import (
	"fmt"
	"os"

	"github.com/daichi-m/go18ds/maps/hashmap"
)

type DrivePaths struct {
	paths *hashmap.Map[string, bool]
}

func New() *DrivePaths {
	var locs DrivePaths

	locs.paths = hashmap.New[string, bool]()
	return &locs
}

func (d *DrivePaths) Load(fileName string) {
	json, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error loading drive paths: ", err)
	}

	d.paths.FromJSON(json)
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

func (d *DrivePaths) IsEnabled(path string) bool {
	value, found := d.paths.Get(path)

	if found {
		return value
	}

	return found // True if the key was found false otherwise. The value variable will be nil.
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
