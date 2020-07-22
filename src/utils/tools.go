package utils

import (
	"crypto/md5"
	"fmt"
	"log"
)

func ErrorHandle(err error, str string) {
	if err != nil {
		log.Fatal(str)
	}
}

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
