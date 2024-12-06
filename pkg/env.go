package env

import (
	"errors"
	"fmt"
	"os"
)

var (
	errMsgEnv    string = "could not load env varliable with key: %s err: %w"
	ErrEnvLookup        = errors.New("structenv-lookupenv")
)

func GetEnv(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf(errMsgEnv, key, ErrEnvLookup)
	}
	return val, nil
}
