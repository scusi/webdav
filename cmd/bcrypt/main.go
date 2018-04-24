// bcrypt - a helper command to hash a password with bcrypt
//
// Usage:
//
// bcrypt -p "my super secret passphrase"
//
package main

import "fmt"
import "flag"
import "log"
import "golang.org/x/crypto/bcrypt"

var passwd string

func init() {
	flag.StringVar(&passwd, "p", "", "password to bcrypt")
}

func main() {
	flag.Parse()
	bpass, err := bcrypt.GenerateFromPassword([]byte(passwd), 8)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("{bcrypt}%s\n", string(bpass))
}
