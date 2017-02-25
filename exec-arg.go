package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	wordPtr := flag.String("name", "Liao Yu", "It is a string.")
	flag.Parse()

	fmt.Println("name:", *wordPtr);

	fmt.Println("pid:", os.Getpid(), "ppid:", os.Getppid())
}
