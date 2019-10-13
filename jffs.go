/*
   Copyright © 2019 M.Watermann, 10247 Berlin, Germany
                  All rights reserved
               EMail : <support@mwat.de>
*/

package jffs

//lint:file-ignore ST1017 - I prefer Yoda conditions

import (
	"net/http"
	"os"
)

/*
 * CREDITS: Brad Fitzpatrick
 * https://groups.google.com/d/msg/golang-nuts/bStLPdIVM6w/hidTJgDZpHcJ
 */

type (
	// Simple struct embedding a `http.File` and ignoring directory reads.
	tNeuteredReaddirFile struct {
		http.File
	}
)

// Readdir reads the contents of the directory associated with file
// and returns a slice of up to `aCount` FileInfo values, as would be
// returned by Lstat, in directory order.
//
// NOTE: This implementation ignores `aCount` and returns nothing, i.e.
// both the FileInfo list and the error are `nil`.
func (f tNeuteredReaddirFile) Readdir(aCount int) ([]os.FileInfo, error) {
	return nil, nil
} // Readdir()

type (
	// Simple struct embedding a `http.FileSystem` that
	// can't read directories.
	justFilesFilesystem struct {
		jffs http.FileSystem
	}
)

// Open is a wrapper around the `Open()` method of the embedded FileSystem
// that returns a `http.File` that can't read directory contents.
func (fs justFilesFilesystem) Open(aName string) (http.File, error) {
	f, err := fs.jffs.Open(aName)
	if nil != err {
		return nil, err
	}

	return tNeuteredReaddirFile{f}, nil
} // Open()

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

// FileServer returns a handler that serves HTTP requests
// with the contents of the file system rooted at `aRoot`.
//
// To use the operating system's file system implementation,
// use http.Dir:
//
//     http.Handle("/", http.FileServer(http.Dir("/tmp")))
//
//	`aRoot` The root of the filesystem to serve.
func FileServer(aRoot http.FileSystem) http.Handler {
	// result.docFS = http.FileServer(http.Dir(CalibreLibraryPath()))
	fs := justFilesFilesystem{aRoot}

	return http.FileServer(fs)
} // FileServer()

/* _EoF_ */
