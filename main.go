package main

import (
	"flag"
	"fmt"
	"github.com/jflude/drumz/sequencer"
	"io"
	"log"
	"os"
)

// Supported DrumKits will insert a constructor in this map on start up.
var knownKits = make(map[string]func() sequencer.DrumKit)

func main() {
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(),
			"Usage:", os.Args[0], "[OPTION...] [TABLATURE-FILE]")
		flag.PrintDefaults()
		fmt.Fprintln(flag.CommandLine.Output(),
			"\nSupported drum kits:-")
		for k, _ := range knownKits {
			fmt.Fprintln(flag.CommandLine.Output(), "  ", k)
		}
	}

	var bpm = flag.Int("bpm", 120, "beats per minute")
	var kind = flag.String("kit", "text", "kind of drum kit")
	flag.Parse()
	if *bpm < 1 {
		log.Fatalln("invalid bpm:", *bpm)
	}
	kit := knownKits[*kind]
	if kit == nil {
		log.Fatalln("invalid kit:", *kind)
	}

	var rdr io.Reader
	name := flag.Arg(0)
	if name == "" {
		name = "(standard input)"
		rdr = os.Stdin
	} else {
		f, err := os.Open(name)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
		rdr = f
	}

	tab, err := sequencer.ReadTab(rdr, kit())
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Name:", name)
	log.Println("Tempo:", *bpm, "bpm")
	sequencer.PlayTab(tab, *bpm, true)
}
