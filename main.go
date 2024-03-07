package main

import (
	"fmt"
	"github.com/asticode/go-astisub"
)

func main() {
	s1, _ := astisub.OpenFile("C:/subtitle/jp/processed/Arakawa Under the Bridge/Arakawa Under the Bridge 01.srt")

	for i := 0; i < len(s1.Items); i++ {
		for j := 0; j < len(s1.Items[i].Lines); j++ {
			fmt.Println(s1.Items[i].Lines[j].String())
		}
	}
}
