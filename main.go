package main

import (
	"fmt"
	"github.com/UMMAN2005/packet"
)

func main() {
	word := packet.Hello()
	word2 := packet.Hi()
	word3 := packet.HowAreYou()
	word4 := packet.WhatIsYourVersion()
	word5 := "That is finish"
	fmt.Println(word, "\n", word2, "\n", word3, "\n", word4, "\n", word5)
	// https://www.reddit.com/r/golang/comments/pqx5tu/am_i_using_go_get_latest_correctly/
}
