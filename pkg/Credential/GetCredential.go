package Credential

import (
	"io/ioutil"
	"strings"
)

//Read secret
func GetCredential(path string) (string, error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(c)), err
}
