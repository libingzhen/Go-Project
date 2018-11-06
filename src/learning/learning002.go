package learning

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The Greeting Object")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
