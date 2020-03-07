/*
   Copyright © 2019, 2020 M.Watermann, 10247 Berlin, Germany
               All rights reserved
           EMail : <support@mwat.de>
*/

package jffs

import (
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestFileServer(t *testing.T) {
	fs1, _ := filepath.Abs("./internal")
	type args struct {
		aRoot string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{" 1", args{fs1}},
		{" 2", args{`/tmp`}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileServer(tt.args.aRoot); nil == got {
				t.Errorf("FileServer() = %v, want (!nul))", got)
				return
			}
		})
	}
} // TestFileServer()

func Test_tJustFilesFilesystem_Open(t *testing.T) {
	dir, _ := filepath.Abs("./internal")
	fs1 := tOnlyFilesFilesystem{http.Dir(dir)}
	type args struct {
		aName string
	}
	tests := []struct {
		name    string
		fs      tOnlyFilesFilesystem
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{" 1", fs1, args{""}, false},
		{" 2", fs1, args{"index.html"}, false},
		{" 3", fs1, args{"internal"}, true},
		{" 4", fs1, args{"does.not.exist"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fs.Open(tt.args.aName)
			if (err != nil) != tt.wantErr {
				t.Errorf("tJustFilesFilesystem.Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (nil == got) && (!tt.wantErr) {
				t.Errorf("tJustFilesFilesystem.Open() = %v, want (!nil)", got)
				return
			}
			if nil != got {
				got.Close()
			}
		})
	}
} // Test_tJustFilesFilesystem_Open()

func Test_tNeuteredReaddirFile_Readdir(t *testing.T) {
	dir, _ := filepath.Abs("./")
	ffs1 := tOnlyFilesFilesystem{http.Dir(dir)}
	nf1, _ := ffs1.Open("")
	var w1 []os.FileInfo
	type args struct {
		aCount int
	}
	tests := []struct {
		name    string
		fields  tNeuteredReaddirFile
		args    args
		want    []os.FileInfo
		wantErr bool
	}{
		// TODO: Add test cases.
		{" 1", tNeuteredReaddirFile{nf1}, args{-1}, w1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tNeuteredReaddirFile{
				File: tt.fields.File,
			}
			got, err := f.Readdir(tt.args.aCount)
			if (err != nil) != tt.wantErr {
				t.Errorf("tNeuteredReaddirFile.Readdir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tNeuteredReaddirFile.Readdir() = %v, want %v", got, tt.want)
			}
		})
	}
} // Test_tNeuteredReaddirFile_Readdir()
