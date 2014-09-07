package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ajmadsen/gough"
)

var (
	infiles stringList
	outfile string
)

type stringList []string

func (s *stringList) String() string {
	return fmt.Sprint(*s)
}

func (s *stringList) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func init() {
	flag.Var(&infiles, "i", "input file, can be given multiple times")
	flag.StringVar(&outfile, "o", "", "output file")
}

func main() {
	flag.Parse()

	if len(infiles) == 0 || outfile == "" {
		flag.Usage()
		os.Exit(1)
	}

	ff.Init()

	inputs := make([]ff.Input, len(infiles))
	for i, s := range infiles {
		ipt := ff.NewInput()
		err := ipt.Open(s)
		if err != nil {
			log.Fatal(err)
		}
		defer ipt.Close()
		ipt.Dump(i)
		for j := 0; j < ipt.NStreams(); j++ {
			ipt.OpenDecoder(j, nil)
		}
		inputs = append(inputs, ipt)
	}
}
