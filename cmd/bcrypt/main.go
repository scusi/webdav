package main

import "fmt"
import "flag"
import "log"
import "golang.org/x/crypto/bcrypt"

//import "encoding/base64"

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

	//str := base64.URLEncoding.EncodeToString(bpass)
	fmt.Printf("{bcrypt}%s\n", string(bpass))

}
