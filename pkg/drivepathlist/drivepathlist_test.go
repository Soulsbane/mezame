package drivepathlist

import (
	"testing"

	"github.com/matryer/is"
)

func TestMain(t *testing.T) {
	is := is.New(t)
	paths := New()

	paths.Add("/home")
	paths.Add("/media")
	is.Equal(paths.NumPaths(), 2)

	is.Equal(paths.IsEnabled("/home"), true)
	paths.Disable("/home")
	is.Equal(paths.IsEnabled("/home"), false)

	paths.Enable("/home")
	is.Equal(paths.IsEnabled("/home"), true)

	is.Equal(paths.IsEnabled("/phonehome"), false) // Not found

	paths.Remove("/home")
	is.Equal(paths.NumPaths(), 1)

	paths.Dump()
}
