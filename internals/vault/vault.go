package vault

import (
	"errors"
	"fmt"
)

var keys = map[string]string{
	"stub": "dsadasjfds8f7y8hsouihfasd",
}

// imitate request to Vault.
// TODO: Replace with real vault request
func GetKey(key string) (string, error) {
	if value, ok := keys[key]; ok {
		return value, nil
	}
	return "", errors.New(fmt.Sprintf("key not found: %s", key))
}
