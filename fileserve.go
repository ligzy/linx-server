package main

import (
	"net/http"
	"os"
	"path"

	"github.com/zenazn/goji/web"
)

func fileServeHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	fileName := c.URLParams["name"]
	filePath := path.Join(Config.filesDir, fileName)
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	// plug file expiry checking here

	http.ServeFile(w, r, filePath)
}
