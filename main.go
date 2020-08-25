package main

import (
	"flag"
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	email := flag.String("email", "", "Email to verify")
	flag.Parse()

	// Validate format.
	if err := checkmail.ValidateFormat(*email); err != nil {
		fmt.Println("ERROR:", err)
		// return
	} else {
		fmt.Println("Format OK")
	}

	// Validate domain.
	if err := checkmail.ValidateHost(*email); err != nil {
		fmt.Println("ERROR:", err)
		// return
	} else {
		fmt.Println("Domain OK")
	}

	// // Validate user.
	// err := checkmail.ValidateHost(*email)
	// if smtpErr, ok := err.(checkmail.SmtpError); ok && err != nil {
	// 	fmt.Printf("ERROR: Code: %s, Msg: %s\n", smtpErr.Code(), smtpErr)
	// } else {
	// 	fmt.Println("User OK")
	// }
}
