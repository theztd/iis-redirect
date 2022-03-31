package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

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

	// Prepare directory structure
	if _, err := os.Stat(filepath.Join(webDirectory, fDir)); err != nil {
		if os.MkdirAll(filepath.Join(webDirectory, fDir), os.ModePerm) != nil {
			log.Panicln("ERR: Unable to create directory", filepath.Join(webDirectory, fDir))
			// return fmt.Errorf("Unable to create directory")
		}
	}

	// Generate script
	if verbose {
		log.Println(" - script:", filepath.Join(webDirectory, fDir, fName)+".ts")
	}

	// script content definition
	script := []byte(fmt.Sprintf("import { REDIRECTS } from './%s_config';\n\nexport async function onRequest(context) {\n  return Response.redirect(REDIRECTS[context.request.url.split('?')[1]], %d);\n}", fName, code))

	// Write content to file
	if err := os.WriteFile(filepath.Join(webDirectory, fDir, fName), script, 0644); err != nil {
		log.Println("ERR:", err)
	}

	// Create redirect config
	if verbose {
		log.Println(" - mapping:", fArgs, " >>> ", to)
	}

	confPath := filepath.Join(webDirectory, fDir, fName+"_config.ts")
	if _, err := os.Stat(confPath); err == nil {
		// File does NOT exists

		// TODO: Append line to file (finish it)
		file, err := os.OpenFile(confPath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Panicln(err)
		}
		defer file.Close()
		if _, err := file.WriteString(fmt.Sprintf("REDIRECTS['%s'] = '%s';\n", from, to)); err != nil {
			log.Panicln(err)
		}
		if verbose {
			log.Println("Config:", confPath, "...Created")
		}

	} else if errors.Is(err, os.ErrNotExist) {
		// File exists
		os.WriteFile(confPath, []byte(fmt.Sprintf("export var REDIRECTS = {};\nREDIRECTS['%s'] = '%s';\n", from, to)), 0644)
		if verbose {
			log.Println("Config:", confPath, "...Appended")
		}
	}

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
		log.Println("ERR:", err, "\n  (detail: ", from, to, code, ")")
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
		log.Println("ERR:", err, "\n  (detail: ", from, to, code, ")")
	}

	// return nil
}
