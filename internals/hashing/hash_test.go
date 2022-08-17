package hashing

import (
	"fmt"
	"github.com/jameskeane/bcrypt"
	"testing"
)

var password = "WyWihatdyd?frub1"
var bad_password = "just a wild guess"

func TestHashing(t *testing.T) {
	// generate a random salt with default rounds of complexity
	salt, _ := bcrypt.Salt()

	// generate a random salt with 10 rounds of complexity
	salt, _ = bcrypt.Salt(10)

	// hash and verify a password with random salt
	hash, _ := bcrypt.Hash(password)
	if bcrypt.Match(password, hash) {
		fmt.Println("They match")
	}

	// hash and verify a password with a static salt
	hash, _ = bcrypt.Hash(password, salt)
	if bcrypt.Match(password, hash) {
		fmt.Println("They match")
	}

	// verify a random password fails to match the hashed password
	if !bcrypt.Match(bad_password, hash) {
		fmt.Println("They don't match")
	}
}
