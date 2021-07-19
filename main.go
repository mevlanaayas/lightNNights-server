package main

import (
	"lightNNights/cmd"
)

func main() {
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
