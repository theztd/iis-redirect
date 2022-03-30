package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func cfParametrizedRedirectOLD(from string, to string, code int) {
	/*
		Deploy script to path depend on from

		mkdir -p $webDirectory/$from - file name
	*/
	// split path to dir and file
	dir, _ := filepath.Split(from)
	if os.MkdirAll(filepath.Join(webDirectory, dir), os.ModePerm) != nil {
		log.Panicln("ERR: Unable to create directory", filepath.Join(webDirectory, dir))
	}

	/*
		generate content in file depend on code
	*/

	// script content definition
	script := []byte(fmt.Sprintf("import { REDIRECTS } from './config';\n\nexport async function onRequest(context) {\n  return Response.redirect(REDIRECTS[context.request.url.split('?')[1]], %d);\n}", code))

	// Write content to file
	if err := os.WriteFile(filepath.Join(webDirectory, dir, "script.js"), script, 0644); err != nil {
		log.Panicln("ERR:", err)
	}
	if verbose {
		log.Println("Create parametrized script:", filepath.Join(webDirectory, dir, "script.js"))
	}

	/*
		create redirect config

		- in case the file already exists, only add new redirect
	*/
	config_path := filepath.Join(webDirectory, dir, "config.js")
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

func cfParametrizedRedirect(from string, to string, code int) {
	/*
		Split from path to directory and file
		- prepare directory structure
		- create FILE_NAME.ts file with redirecting script
		- create file FILE_NAME-config.ts and append redirects
	*/
	ret := strings.Split(from, "?")
	fDir, fName := filepath.Split(ret[0])
	fArgs := ret[1]

	if verbose {
		log.Println(" - script:", filepath.Join(webDirectory, fDir, fName)+".ts")
	}

	if verbose {
		log.Println(" - mapping:", fArgs, " >>> ", to)
	}

	// create directory structure

	/*



		Pokracovat zde



	*/

}

func cfFileRedirect(from string, to string, code int) {
	/*
		Split from path to directory and file.
		Than prepare directory path and pass redirection to file
	*/
	if verbose {
		log.Println(" - script:", filepath.Join(webDirectory, from))
	}

	// Prepare directory structure
	fDir, _ := filepath.Split(from)
	if _, err := os.Stat(filepath.Join(webDirectory, fDir)); err != nil {
		if os.MkdirAll(filepath.Join(webDirectory, fDir), os.ModePerm) != nil {
			log.Panicln("ERR: Unable to create directory", filepath.Join(webDirectory, fDir))
			// return fmt.Errorf("Unable to create directory")
		}
	}

	// Create script
	script := []byte(fmt.Sprintf("export async function onRequest(context) {\n  return Response.redirect(%s, %d);\n}", to, code))

	if err := os.WriteFile(filepath.Join(webDirectory, from), script, 0644); err != nil {
		log.Fatalln("ERR:", err, "\n  (detail: ", from, to, code, ")")
	}

}

func cfSimpleRedirect(from string, to string, code int) { //(err error) {
	/*
		Prepare directory structure and create index.ts with redirecting code
	*/
	if verbose {
		log.Println(" - script:", filepath.Join(webDirectory, from, "index.ts"))
	}

	// Prepare directory structure
	if _, err := os.Stat(filepath.Join(webDirectory, from)); err != nil {
		if os.MkdirAll(filepath.Join(webDirectory, from), os.ModePerm) != nil {
			log.Panicln("ERR: Unable to create directory", filepath.Join(webDirectory, from))
			// return fmt.Errorf("Unable to create directory")
		}
	}

	//Create script
	script := []byte(fmt.Sprintf("export async function onRequest(context) {\n  return Response.redirect(%s, %d);\n}", to, code))

	if err := os.WriteFile(filepath.Join(webDirectory, from, "index.ts"), script, 0644); err != nil {
		log.Fatalln("ERR:", err, "\n  (detail: ", from, to, code, ")")
	}

	// return nil
}
