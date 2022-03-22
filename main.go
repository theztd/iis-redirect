package main

import (
	"flag"
	"log"
)

var web_directory, redirects_file string
var verbose bool

func main() {
	// CLI interface
	flag.StringVar(&redirects_file, "i", "rewriteMap.config", "Path to file for parsing (IIS rewriteMap.config).")
	flag.StringVar(&web_directory, "o", "./tmp_functions", "Path to directory, where redirects will be placed.")
	flag.BoolVar(&verbose, "v", false, "Verbose output.")
	flag.Parse()

	rewrites, err := parseRewriteMap(redirects_file)
	if err != nil {
		log.Panicln("No data from parsing")
	}
	for _, r := range rewrites {
		createCfRedirect(r.From, r.To, r.Type)
	}
}
