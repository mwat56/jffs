package main

//lint:file-ignore ST1017 - I prefer Yoda conditions

import (
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/mwat56/jffs"
)

func main() {
	dir, _ := filepath.Abs("./internal/")
	server := &http.Server{
		Addr:              ":4321",
		Handler:           jffs.FileServer(http.Dir(dir)),
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	if err := server.ListenAndServe(); nil != err {
		log.Fatal(err)
	}
} // main()

/* _EoF_ */
