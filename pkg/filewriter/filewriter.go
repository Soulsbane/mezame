package filewriter

import "github.com/Soulsbane/mezame/pkg/drivepaths"

const (
	DefaultMezameFileName = ".mezame"
)

type FileWriter struct {
	paths *drivepaths.DrivePaths
}

func New() *FileWriter {
	var writer FileWriter
	writer.paths = drivepaths.New()

	return &writer
}

func (f *FileWriter) Load(fileName string) {
	f.paths.Load(fileName)
}
