package locations

import (
	"fmt"

	"github.com/daichi-m/go18ds/maps/hashmap"
)

type Locations struct {
	pathsFileName string
	paths         *hashmap.Map[string, bool]
}

func New(fileName string) *Locations {
	var locs Locations

	locs.paths = hashmap.New[string, bool]()
	locs.pathsFileName = fileName

	return &locs
}

func (f *Locations) AddPath(path string) {
	f.paths.Put(path, true)
}

func (f *Locations) RemovePath(path string) {
	f.paths.Remove(path)
}

func (f *Locations) EnablePath(path string) {
	f.paths.Put(path, true)
}

func (f *Locations) DisablePath(path string) {
	f.paths.Put(path, false)
}

func (f *Locations) Dump() {
	for _, key := range f.paths.Keys() {
		value, _ := f.paths.Get(key)
		fmt.Println(key, value)
	}
}
