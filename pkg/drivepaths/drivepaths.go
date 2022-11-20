package drivepaths

import (
	"encoding/json"
	"fmt"
	"os"
)

type DrivePaths struct {
	paths map[string]bool
}

func New() *DrivePaths {
	var locs DrivePaths

	locs.paths = make(map[string]bool)
	return &locs
}

func (d *DrivePaths) Load(fileName string) {
	fileData, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error loading drive paths file: ", err)
	}

	err = json.Unmarshal(fileData, &d.paths)

	if err != nil {
		fmt.Println("Failed to unmarshal json: ", err)
	}
}

func (d *DrivePaths) Save(fileName string) {
	data, err := json.Marshal(d.paths)

	if err != nil {
		fmt.Println("Failed to marshal json: ", err)
	} else {
		err = os.WriteFile(fileName, data, 0644)

		if err != nil {
			fmt.Println("Error saving to drive paths: ", err)
		}
	}
}

func (d *DrivePaths) Add(path string) {
	d.paths[path] = true
}

func (d *DrivePaths) Remove(path string) {
	delete(d.paths, path)
}

func (d *DrivePaths) Enable(path string) {
	d.paths[path] = true
}

func (d *DrivePaths) Disable(path string) {
	d.paths[path] = false
}

func (d *DrivePaths) IsEnabled(path string) bool {
	if _, ok := d.paths[path]; ok {
		return d.paths[path]
	} else {
		return false
	}

}

func (d *DrivePaths) NumPaths() int {
	return len(d.paths)
}

func (d *DrivePaths) Dump() {
	for path, enabled := range d.paths {
		fmt.Printf("%s => %t\n", path, enabled)
	}
}
