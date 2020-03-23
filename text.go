package main

import (
	"fmt"
	"github.com/jflude/drumz/sequencer"
)

var knownTextDrums = map[string]bool{
	"CY": true, // cymbal
	"HH": true, // hi-hat
	"TA": true, // tambourine
	"CW": true, // cowbell
	"HT": true, // high tom
	"MT": true, // mid tom
	"LT": true, // low tom
	"SD": true, // snare drum
	"BD": true, // bass drum
	"AC": true, // accent
}

func init() {
	knownKits["text"] =
		func() sequencer.DrumKit { return textKit{} }
}

type textKit struct{}

func (tk textKit) HasDrum(name string) bool {
	return knownTextDrums[name]
}

func (tk textKit) NewLine() sequencer.TabLine {
	return new(textLine)
}

type textLine struct {
	val string
}

func (tl *textLine) AddDrum(name string) {
	if tl.val != "" {
		tl.val += "+"
	}
	tl.val += name
}

func (tl *textLine) Play(lastLine bool) {
	if tl.val == "" {
		fmt.Print("__")
	} else {
		fmt.Print(tl.val)
	}
	if lastLine {
		fmt.Println()
	}
}
