package main

import "os"

func main() {
	for i, v := range os.Args{
		if i != 0 {
		os.Stdout.WriteString(v)
		os.Stdout.WriteString("\n")
		}
	}
}