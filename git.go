package gosense

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func has_git(path string) bool {
	if path[len(path)-1:] != "/" {
		path += "/"
	}

	if _, err := os.Stat(path + ".git"); err == nil {
		return true
	}

	return false
}

func is_local_git(path string) bool {
	if !has_git(path) {
		return false
	}

	if path[len(path)-1:] != "/" {
		path += "/"
	}

	contents, err := ioutil.ReadFile(path + ".git/config")
	if err != nil {
		fmt.Println(err)
	}

	if !strings.Contains(string(contents), "url") {
		return false
	}

	return true
}
