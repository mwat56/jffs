package jffs

//lint:file-ignore ST1017 - I prefer Yoda conditions

import (
	"net/http"
	"os"
	"strings"
)

/*
 * CREDITS:
 *
 * Brad Fitzpatrick
 * https://groups.google.com/d/msg/golang-nuts/bStLPdIVM6w/hidTJgDZpHcJ
 *
 * Alex Edwards
 * https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings
 */

type (
	// Simple struct embedding a `http.File` and ignoring
	// directory reads.
	tNeuteredReaddirFile struct {
		http.File
	}

	// Simple struct embedding a `http.FileSystem` that
	// can't read directories.
	tOnlyFilesFilesystem struct {
		offs http.FileSystem
	}
)

// Readdir reads the contents of the directory associated with file
// and returns a slice of up to `aCount` FileInfo values, as would
// be returned by Lstat, in directory order.
//
// NOTE: This implementation ignores `aCount` and returns nothing,
// i.e. both the `FileInfo` list and the `error` are `nil`.
func (f tNeuteredReaddirFile) Readdir(aCount int) ([]os.FileInfo, error) {
	return nil, nil
} // Readdir()

// Open is a wrapper around the `Open()` method of the embedded FileSystem
// that returns a `http.File` which can't read directory contents.
//
//	`aName` is the name of the file to open.
func (ffs tOnlyFilesFilesystem) Open(aName string) (http.File, error) {
	f, err := ffs.offs.Open(aName)
	if nil != err {
		return nil, err
	}

	if fInfo, err := f.Stat(); (nil == err) && fInfo.IsDir() {
		f.Close() // free directory's file handle

		idx := strings.TrimSuffix(aName, `/`) + `/index.html`
		if f, err = ffs.offs.Open(idx); nil != err {
			return nil, err
		}
	}

	return tNeuteredReaddirFile{f}, nil
} // Open()

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

// FileServer returns a handler that serves HTTP requests
// with the contents of the file system rooted at `aRoot`.
//
// To use the operating system's file system implementation,
// use `http.Dir()`:
//
//	myHandler := http.FileServer(http.Dir("/tmp")))
//
// To use this implementation you'd use:
//
//	myHandler := jffs.FileServer(http.Dir("/tmp")))
//
//	`aRoot` The root of the filesystem to serve.
func FileServer(aRoot http.FileSystem) http.Handler {
	return http.FileServer(tOnlyFilesFilesystem{aRoot})
} // FileServer()

/* _EoF_ */
