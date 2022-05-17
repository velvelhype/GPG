package main

import "os"
// import "fmt"
func main() {
	args := os.Args[1:]
	count := 0
	for range args{
		count++
	}
	for i := range args{
		for j := i + 1; j < count; j++ {
		if args[i][0] > args[j][0]{
			args[i], args[j] = args[j], args[i]
		}
		}
	}
	for _, v := range args {
		os.Stdout.WriteString(v)
		os.Stdout.WriteString("\n")
	}
}