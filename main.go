package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/prometheus/log"
)

var outName = flag.String("out", "-", "Target for writing the output")

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [args ...] <config_file>", flag.Arg(0))

	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Parse()

	var (
		err error
		in  io.Reader = os.Stdin
		out io.Writer = os.Stdout
	)

	if flag.NArg() > 0 {
		filename := flag.Args()[0]
		in, err = os.Open(filename)
		if err != nil {
			log.Fatalf("error opening input file: %s", err)
		}
		log.Infof("translating file %s", filename)
	}

	translate(in, out)
}

func translate(in io.Reader, out io.Writer) error {
	return nil
}
