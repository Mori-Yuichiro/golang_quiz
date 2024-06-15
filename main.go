package main

import (
	"fmt"

	"github.com/yourname/reponame/quiz"
)

func main() {
	m := map[int]string{
		1: "01",
		2: "02",
		3: "03",
	}

	key, err := quiz.FindKeyByValue(m, "03")
	fmt.Println(key, err)
}
