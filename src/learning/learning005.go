package learning

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
        flag.PrintDefaults()
        }
	flag.StringVar(&name, "name", "everyone", "The Greeting Object")
}

func main() {
	flag.Parse()
	hello(name)
}
