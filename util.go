package sms

import (
	"crypto/md5"
	"fmt"
)

func MapKeys(v map[string]string) []string {

	keys := []string{}
	for k := range v {
		keys = append(keys, k)
	}
	return keys
}

func GetMd5String(content []byte) string {
	password := md5.Sum(content)
	return fmt.Sprintf("%x", password)
}
