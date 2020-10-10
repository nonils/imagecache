package utils

import (
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"net/http"
	"os"
)

func CheckStatusCodeIs200(res *http.Response) error {
	if res.StatusCode > 299 || res.StatusCode < 200 {
		return errors.New("the response was not sucess")
	}
	return nil
}

func CheckError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		fmt.Println(err)
	}
}
func GetConfigValueFromKey(key string) string {
	return os.Getenv(key)
}

var Cache *cache.Cache
