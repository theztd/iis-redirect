package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func createCfRedirect(from string, to string, code int) {
	/*
		Deploy script to path depend on from

		mkdir -p $redirects_directory/$from - file name
	*/
	// split path to dir and file
	dir, fName := filepath.Split(from)
	if os.MkdirAll(filepath.Join(redirects_directory, dir), os.ModePerm) != nil {
		panic("Unable to create directory")
	}

	/*
		generate content in file depend on code
	*/

	// script content definition
	script := []byte(fmt.Sprintf("import { REDIRECTS } from './config';\n\nexport async function onRequest(context) {\n  return Response.redirect(REDIRECTS[context.request.url.split('?')[1]], %d);\n}", code))

	// Write content to file
	if os.WriteFile(filepath.Join(redirects_directory, dir, "script.js"), script, 0644) != nil {
		panic(fmt.Sprintf("Unable to deploy script"))
	}

	fmt.Println(fName)

	/*
		create redirect config

		- in case the file already exists, only add new redirect
	*/
	config_path := filepath.Join(redirects_directory, dir, "config.js")
	if _, err := os.Stat(config_path); err == nil {
		// soubor uz existuje
		fmt.Println("Exists")
	} else if errors.Is(err, os.ErrNotExist) {
		// Soubor nexistuje
		os.WriteFile(config_path, []byte(fmt.Sprintf("export var REDIRECTS = {};\nREDIRECTS['%s'] = '%s';\n}", from, to)), 0644)
	}

}
