package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

/*
	Example of parsing structure

	<?xml version="1.0" encoding="utf-8"?>
	<rewriteMaps>
	  <rewriteMap name="PermanentRedirects" defaultValue="">
		<add key="/from1" value="dest1" />
	  </rewriteMap>
	</rewriteMaps>
*/
type rewriteMaps struct {
	Rewrites []rewriteMap `xml:"rewriteMap"`
}

type rewriteMap struct {
	Name     string    `xml:"name,attr"`
	Rewrites []rewrite `xml:"add"`
}

type rewrite struct {
	From string `xml:"key,attr"`
	To   string `xml:"value,attr"`
}

// END: Parsing input structure

// Return from parseRewriteMap
type rewrites struct {
	From string
	To   string
	Type int
}

func parseRewriteMap(file_path string) (redirects []rewrites, err error) {
	redirectType := make(map[string]int)

	redirectType["PermanentRedirects"] = 301
	redirectType["TempRedirects"] = 307

	var ret []rewrites

	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("Unable to open file")
		return ret, err
	}
	defer file.Close()

	bValue, _ := ioutil.ReadAll(file)

	var out rewriteMaps
	if xml.Unmarshal(bValue, &out) != nil {
		fmt.Println("Unable to parse file", file_path)
	}

	for _, k := range out.Rewrites {
		code := redirectType[k.Name]
		for _, r := range k.Rewrites {
			ret = append(ret, rewrites{From: r.From, To: r.To, Type: code})
		}

	}

	return ret, nil
}
