package drivepaths

import (
	"testing"

	"github.com/matryer/is"
)

func TestMain(t *testing.T) {
	is := is.New(t)
	paths := New("drivepaths.json")

	paths.Add("/home")
	paths.Add("/media")
	is.Equal(paths.NumPaths(), 2)

	paths.Remove("/home")
	is.Equal(paths.NumPaths(), 1)
}
