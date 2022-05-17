package main

import "os"

func main() {
	count := 0
	for range os.Args{
		count++
	}
	for i := count; i > 1; i--{
		os.Stdout.WriteString(os.Args[i - 1])
		os.Stdout.WriteString("\n")
	}
}