package filewriter

import drivepathlist "github.com/Soulsbane/mezame/pkg/drivepathlist"

const (
	DefaultMezameFileName = ".mezame"
)

type FileWriter struct {
	paths *drivepathlist.DrivePathList
}

func New() *FileWriter {
	var writer FileWriter
	writer.paths = drivepathlist.New()

	return &writer
}

func (f *FileWriter) Load(fileName string) {
	f.paths.Load(fileName)
}
