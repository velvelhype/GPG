package main

import "os"

func main() {
	thrash_pos := 0
	for i,v := range os.Args[0]{
		if v == '/' {
			thrash_pos = i + 1
		}
	}
	os.Stdout.WriteString(os.Args[0][thrash_pos:])
	os.Stdout.WriteString("\n")
}