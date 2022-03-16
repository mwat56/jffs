/*
   Copyright © 2019, 2022 M.Watermann, 10247 Berlin, Germany
                  All rights reserved
               EMail : <support@mwat.de>
*/
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
		Handler:           jffs.FileServer(dir),
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
} // main()

/* _EoF_ */
