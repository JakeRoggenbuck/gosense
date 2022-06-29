package main

import (
	"os"
)

type Lang int

const (
	Python Lang = iota + 1
	Rust
	Java
	JS
	Go
	None
)

func (l Lang) String() string {
	langs := [...]string{"Python", "Rust", "Java", "JavaScript", "Go", "None"}
	return langs[l-1]
}

func has_sub_file(path string, sub string) bool {
	if _, err := os.Stat(path + sub); err == nil {
		return true
	}

	return false
}

func get_lang(path string) Lang {
	if path[len(path)-1:] != "/" {
		path += "/"
	}

	if has_sub_file(path, "setup.py") {
		return Python
	} else if has_sub_file(path, "package.json") {
		return JS
	} else if has_sub_file(path, "Cargo.toml") {
		return Rust
	} else if has_sub_file(path, "pom.xml") {
		return Java
	} else if has_sub_file(path, "go.mod") {
		return Go
	}

	return None
}
