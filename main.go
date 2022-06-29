package main

import (
	"fmt"
)

func main() {
	git := has_git("./")
	if git {
		fmt.Println("Is git!")
	} else {
		fmt.Println("Is not git.")
	}

	local := is_local_git("./")
	if local {
		fmt.Println("Is local!")
	} else {
		fmt.Println("Is not local.")
	}
}
