package main

import (
	// "bufio"
	"fmt"
	// "io"
	"flag"
	"os"
)

func main() {
	songCmd := flag.NewFlagSet("song", flag.ExitOnError)

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Error: Try --help for more information")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("Error: Too many arguments")
		os.Exit(1)
	}

	switch args[0] {

	case "song":
		songCmd.Parse(args[1:])
		fmt.Println("Song Info:")
	case "play":
		fmt.Println("Song played")
	case "pause":
		fmt.Println("Song paused")
	}
}
