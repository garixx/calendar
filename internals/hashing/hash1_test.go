package hashing

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestBcryptingIsEasy(t *testing.T) {
	pass := []byte("mypassword")
	hp, err := bcrypt.GenerateFromPassword(pass, 0)
	fmt.Println("hp=", string(hp))
	if err != nil {
		t.Fatalf("GenerateFromPassword error: %s", err)
	}

	if bcrypt.CompareHashAndPassword(hp, pass) != nil {
		t.Errorf("%v should hash %s correctly", hp, pass)
	}

	notPass := "notthepass"
	err = bcrypt.CompareHashAndPassword(hp, []byte(notPass))
	if err != bcrypt.ErrMismatchedHashAndPassword {
		t.Errorf("%v and %s should be mismatched", hp, notPass)
	}
}

func TestBcryptingIsEasy2(t *testing.T) {

	pass := []byte("password")
	hp := []byte("$2a$14$1MEzFbcJXlcBQ8/26tTwK.yF7A2k3TcjWhetZ3CZlvNxUz5LX/NwG")

	if bcrypt.CompareHashAndPassword(hp, pass) != nil {
		t.Errorf("%v should hash %s correctly", hp, pass)
	}
}
