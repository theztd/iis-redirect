package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func createCfRedirect(from string, to string, code int) {
	/*
		Deploy script to path depend on from

		mkdir -p $web_directory/$from - file name
	*/
	// split path to dir and file
	dir, _ := filepath.Split(from)
	if os.MkdirAll(filepath.Join(web_directory, dir), os.ModePerm) != nil {
		log.Panicln("ERR: Unable to create directory", filepath.Join(web_directory, dir))
	}

	/*
		generate content in file depend on code
	*/

	// script content definition
	script := []byte(fmt.Sprintf("import { REDIRECTS } from './config';\n\nexport async function onRequest(context) {\n  return Response.redirect(REDIRECTS[context.request.url.split('?')[1]], %d);\n}", code))

	// Write content to file
	if err := os.WriteFile(filepath.Join(web_directory, dir, "script.js"), script, 0644); err != nil {
		log.Panicln("Unable to deploy script")
	}
	if verbose {
		log.Println("Create script:", filepath.Join(web_directory, dir, "script.js"))
	}

	/*
		create redirect config

		- in case the file already exists, only add new redirect
	*/
	config_path := filepath.Join(web_directory, dir, "config.js")
	if _, err := os.Stat(config_path); err == nil {
		// File does NOT exists

		// TODO: Append line to file (finish it)
		file, err := os.OpenFile(config_path, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		if _, err := file.WriteString(fmt.Sprintf("REDIRECTS['%s'] = '%s';\n", from, to)); err != nil {
			panic(err)
		}
		if verbose {
			log.Println("Config:", config_path, "...Created")
		}

	} else if errors.Is(err, os.ErrNotExist) {
		// File exists
		os.WriteFile(config_path, []byte(fmt.Sprintf("export var REDIRECTS = {};\nREDIRECTS['%s'] = '%s';\n", from, to)), 0644)
		if verbose {
			log.Println("Config:", config_path, "...Appended")
		}
	}

}
