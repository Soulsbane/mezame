package drivepathlist

import (
	"encoding/json"
	"fmt"
	"os"
)

type DrivePathList struct {
	paths map[string]bool
}

func New() *DrivePathList {
	var locs DrivePathList

	locs.paths = make(map[string]bool)
	return &locs
}

func (d *DrivePathList) Load(fileName string) error {
	fileData, err := os.ReadFile(fileName)

	if err != nil {
		return err
	} else {
		return json.Unmarshal(fileData, &d.paths)
	}
}

func (d *DrivePathList) Save(fileName string) error {
	data, err := json.Marshal(d.paths)

	if err != nil {
		return err
	} else {
		return os.WriteFile(fileName, data, 0644)
	}
}

func (d *DrivePathList) Add(path string) {
	d.paths[path] = true
}

func (d *DrivePathList) Remove(path string) {
	delete(d.paths, path)
}

func (d *DrivePathList) Enable(path string) {
	d.paths[path] = true
}

func (d *DrivePathList) Disable(path string) {
	d.paths[path] = false
}

func (d *DrivePathList) IsEnabled(path string) bool {
	if _, ok := d.paths[path]; ok {
		return d.paths[path]
	} else {
		return false
	}
}

func (d *DrivePathList) Length() int {
	return len(d.paths)
}

func (d *DrivePathList) Dump() {
	for path, enabled := range d.paths {
		fmt.Printf("%s => %t\n", path, enabled)
	}
}
