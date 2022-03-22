package main

import (
	"flag"
	"fmt"
)

var redirects_file string

func main() {
	flag.StringVar(&redirects_file, "f", "", "Path to file for parsing")
	flag.Parse()

	rewrites, err := parseRewriteMap(redirects_file)
	if err != nil {
		fmt.Println("No data from parsing")
	}
	for _, r := range rewrites {
		fmt.Println(">>>", r.From, r.To, r.Type)
		createCfRedirect(r.From, r.To, r.Type)
	}
}
