package main

import (
	"flag"
)

func main() {
	tool := flag.String("tool", "verifyemail", "verifyemail, mxlookup")
	email := flag.String("email", "", "Email to verify")
	domain := flag.String("domain", "", "Domain to verify")
	flag.Parse()

	switch *tool {
	case "verifyemail":
		verifyEmail(*email)
	case "mxlookup":
		mxLookup(*domain)
	default:
		verifyEmail(*email)
	}
}
