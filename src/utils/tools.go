package utils

import (
	"crypto/md5"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func ErrorHandle(err error, str string) {
	if err != nil {
		log.Fatal(str)
	}
}

//返回字符串的MD5
func Md5(buf string) string {
	hash := md5.New()
	hash.Write([]byte(buf))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//返回字符串的随机加密hash
func GeneratePassword(str string) (result string, err error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	result = string(pwd)
	return
}

//将明文字符串和随机加密hash比较, 验证有效性
func VaildataPassword(pwd, hashPwd string) (result bool, err error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd)); err == nil {
		result = true
	} else {
		err = errors.New("verification failed.")
	}
	return
}
