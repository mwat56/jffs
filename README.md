# jffs – Just Files Filesystem

[![golang](https://img.shields.io/badge/Language-Go-green.svg)](https://golang.org)
[![GoDoc](https://godoc.org/github.com/mwat56/jffs?status.svg)](https://godoc.org/github.com/mwat56/jffs)
[![Go Report](https://goreportcard.com/badge/github.com/mwat56/jffs)](https://goreportcard.com/report/github.com/mwat56/jffs)
[![Issues](https://img.shields.io/github/issues/mwat56/jffs.svg)](https://github.com/mwat56/jffs/issues?q=is%3Aopen+is%3Aissue)
[![Size](https://img.shields.io/github/repo-size/mwat56/jffs.svg)](https://github.com/mwat56/jffs/)
[![Tag](https://img.shields.io/github/tag/mwat56/jffs.svg)](https://github.com/mwat56/jffs/tags)
[![License](https://img.shields.io/github/mwat56/jffs.svg)](https://github.com/mwat56/jffs/blob/master/LICENSE)

- [jffs – Just Files Filesystem](#jffs-%e2%80%93-just-files-filesystem)
	- [Purpose](#purpose)
	- [Installation](#installation)
	- [Usage](#usage)
	- [Licence](#licence)

----

## Purpose

Often you don't want remote users to access your filesystem's static files outside the URLs you provide with your web-application.
Especially if there are file(s) in your (sub-)directories which shouldn't be served to others.
Therefore it's imperative to hamper the Go fileserver insofar as to _not_ produce directory listings showing all available files.
That's were this small package comes in.

> Credits go to [Brad Fitzpatrick](https://groups.google.com/d/msg/golang-nuts/bStLPdIVM6w/hidTJgDZpHcJ) and [Alex Edwards](https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings).

## Installation

You can use `Go` to install this package for you:

    go get -u github.com/mwat56/jffs

## Usage

This package exports a single function `FileServer()`.
So, while you're used to call

	myStaticDirectory :="/tmp"
	myStaticHandler := http.FileServer(http.Dir(myStaticDirectory)))

to use _this_ implementation you'd now just use:

	myStaticHandler := jffs.FileServer(myStaticDirectory)

That's all.

## Licence

    Copyright © 2019, 2020 M.Watermann, 10247 Berlin, Germany
                    All rights reserved
                EMail : <support@mwat.de>

> This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.
>
> This software is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
>
> You should have received a copy of the GNU General Public License along with this program. If not, see the [GNU General Public License](http://www.gnu.org/licenses/gpl.html) for details.

----
