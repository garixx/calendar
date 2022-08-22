package ssm

import (
	"errors"
	"fmt"
)

var secrets = map[string]string{
	"stub": "dsadasjfds8f7y8hsouihfasd",
}

// GetSecret imitate request to Secrets store (Vault, for example).
// TODO: Replace with real ssm request
func GetSecret(secret string) (string, error) {
	if value, ok := secrets[secret]; ok {
		return value, nil
	}
	return "", errors.New(fmt.Sprintf("key not found: %s", secret))
}
