package locations

type Locations struct {
	pathsFileName string
	paths         []string
}

func New(fileName string) *Locations {
	var locations Locations

	locations.pathsFileName = fileName
	return &locations
}
