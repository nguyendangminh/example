package main 

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Fprint(os.Stdout, "minhnd\n")
}