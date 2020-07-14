package utils

import "log"

func ErrorHandle(err error, str string) {
	if err != nil {
		log.Fatal(str)
	}
}
