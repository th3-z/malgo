package utils

import (
    "io/ioutil"
)

func FileToStr (filename string) string {
	fBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(fBytes)
}

