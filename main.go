package main

import (
	"fmt"
	"github.com/UMMAN2005/packet"
)

func main() {
	word := packet.Hello()
	word2 := packet.Hi()
	word3 := packet.How_are_you()
	fmt.Println(word, "\n", word2, "\n", word3)
}
