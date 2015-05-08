package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/prometheus/log"

	"github.com/prometheus/migrate/v0x13"
)

var outName = flag.String("out", "-", "Target for writing the output")

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [args ...] [<config_file>]", flag.Arg(0))

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
			log.Fatalf("Error opening input file: %s", err)
		}
		log.Infof("Translating file %s", filename)
	}

	if err := translate(in, out); err != nil {
		log.Fatal(err)
	}
}

func translate(in io.Reader, out io.Writer) error {
	b, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	var oldConfig v0x13.Config
	err = proto.UnmarshalText(string(b), &oldConfig.PrometheusConfig)
	if err != nil {
		return fmt.Errorf("Error parsing old config file: %s", err)
	}

	fmt.Println(oldConfig)
	return nil
}
