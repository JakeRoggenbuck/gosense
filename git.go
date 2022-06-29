package main

import (
	"os"
	"fmt"
)

func has_git(path string) bool {
	if _, err := os.Stat(path + "/.git"); err == nil {
		return true
	}

	return false
}

func is_local_git(path string) bool {
	if !has_git(path) {
		return false
	}

	f, err := os.Open(path + "/.git/config")
	if err != nil {
		fmt.Println(err)
	}
	
	defer f.Close()



	return true
}
