package main

import (
	"flag"
	"log"
	"strings"
)

var webDirectory, redirectsFile string
var verbose bool

func main() {
	// CLI interface
	flag.StringVar(&redirectsFile, "i", "rewriteMap.config", "Path to file for parsing (IIS rewriteMap.config).")
	flag.StringVar(&webDirectory, "o", "./tmp_functions", "Path to directory, where redirects will be placed.")
	flag.BoolVar(&verbose, "v", false, "Verbose output.")
	flag.Parse()

	rewrites, err := parseRewriteMap(redirectsFile)
	if err != nil {
		flag.Usage()
		log.Fatalln("Nothing to parse")
	}

	for _, r := range rewrites {

		if strings.Contains(r.From, ".") && strings.Contains(r.From, "?") == false {
			log.Println("File:", r.From)
			cfFileRedirect(r.From, r.To, r.Type)
			continue
		}

		if strings.Contains(r.From, "?") {
			log.Println("Parametrized:", r.From)
			cfParametrizedRedirect(r.From, r.To, r.Type)
			continue
		}

		log.Println("Simple:", r.From)
		cfSimpleRedirect(r.From, r.To, r.Type)

	}
}
