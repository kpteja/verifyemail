package main

import (
	"fmt"
	"net"
)

func mxLookup(domain string) {
	mxs, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	if len(mxs) > 0 {
		fmt.Println("Found:")
		for _, mx := range mxs {
			fmt.Printf("%s %v\n", mx.Host, mx.Pref)
		}
	} else {
		fmt.Println("No records found")
	}
}
