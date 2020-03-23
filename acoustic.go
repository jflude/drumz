package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/jflude/drumz/sequencer"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Needs to be short to ensure playback starts "immediately".
const bufferTime = time.Second / 20

var knownAcousticDrums = map[string]bool{
	"HH": true, // hi-hat
	"SD": true, // snare drum
	"BD": true, // bass drum
}

func init() {
	knownKits["acoustic"] =
		func() sequencer.DrumKit { return newAcousticKit() }
}

type acousticKit struct {
	samples map[string]beep.Streamer
}

func newAcousticKit() *acousticKit {
	kit := &acousticKit{make(map[string]beep.Streamer)}
	var format beep.Format
	for name, _ := range knownAcousticDrums {
		p := filepath.Join("samples", strings.ToLower(name)+".wav")
		f, err := os.Open(p)
		if err != nil {
			panic(err)
		}
		kit.samples[name], format, err = wav.Decode(f)
		if err != nil {
			panic(err)
		}
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(bufferTime))
	return kit
}

func (ak *acousticKit) HasDrum(name string) bool {
	return knownAcousticDrums[name]
}

func (ak *acousticKit) NewLine() sequencer.TabLine {
	return &acousticLine{ak, nil}
}

type acousticLine struct {
	kit       *acousticKit
	streamers []beep.Streamer
}

func (al *acousticLine) AddDrum(name string) {
	al.streamers = append(al.streamers, al.kit.samples[name])
}

func (al *acousticLine) Play(lastLine bool) {
	if len(al.streamers) == 0 {
		return
	}
	for i := 0; i < len(al.streamers); i++ {
		al.streamers[i].(beep.StreamSeeker).Seek(0)
	}
	// Play the samples without waiting for them to all complete, as they
	// are of various lengths and we don't want to limit the tempo.
	speaker.Play(beep.Seq(beep.Mix(al.streamers...)))
}
