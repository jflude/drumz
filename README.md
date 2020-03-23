# Splice Interview Exercise - Justin Flude - Drum Machine

## How to Test and Build

```
$ go test ./...
$ go build
```

On MacOS Catalina (10.15) this will output a number of deprecation warnings
for the audio library used, which can be safely ignored.  Please note that
the program has not been built or tested on any other platform.

## How to Run

```
$ ./drumz -help
Usage: ./drumz [OPTION...] [TABLATURE-FILE]
  -bpm int
    	beats per minute (default 120)
  -kit string
    	kind of drum kit (default "text")

Supported drum kits:-
   acoustic
   text
```

## Example of Output

```
$ ./drumz examples/given.txt 
Name: examples/given.txt
Tempo: 120 bpm
BD__HH__SD+BD__HH__BD__HH__SD+BD__HH__
BD__HH__SD+BD__HH__BD__HH__SD+BD__HH__
BD__HH__SD+BD__HH__BD__HH__SD+BD__HH__
BD__HH__SD+BD__HH__BD__HH__SD+BD__HH__
BD__HH__SD+BD__HH__^C
```

### Tablature Format

This is described in the sequencer documentation for the ReadTab function, at
the end of this file.

## Provided Drum Tablatures

In the examples/ directory are some drum tablatures to try out:-

    The Ramones, "Blitzkrieg Bop"
    James Brown, "Funky Drummer"
    Led Zeppelin, "When the Levee Breaks"
    Splice, "The Given Example"

## Extending the Program

New drum "kits" which output in different ways can be added to the program by
implementing the DrumKit and TabLine interfaces found in the sequencer package,
registering the new kit on startup with the knownKits map (see text.go and
acoustic.go)

New drum types (sounds) can be added to the acoustic kit by adding WAV files
to the samples/ directory and updating knownAcousticDrums in acoustic.go

## Documentation

```
$ go doc -all sequencer
package sequencer // import "github.com/jflude/drumz/sequencer"

Package sequencer provides interfaces and functions to read and play drum
tablatures.

CONSTANTS

const LinesPerBeat = 4
    LinesPerBeat determines how many lines in a tablature correspond to a single
    beat.


VARIABLES

var ErrInvalidTab = errors.New("invalid tablature")
    ErrInvalidTab is returned by ReadTab for tablatures with incorrect formats.

var ErrUnknownDrum = errors.New("unknown drum")
    ErrUnknownDrum is returned by ReadTab for tablatures with drums which
    DrumKit does not support or recognise.


FUNCTIONS

func PlayTab(tab DrumTab, bpm int, repeat bool)
    PlayTab takes a DrumTab created by ReadTab and plays it, in some way, at the
    tempo given by the bpm argument. If repeat is false it will be played just
    once.


TYPES

type DrumKit interface {
	HasDrum(name string) bool
	NewLine() TabLine
}
    DrumKit is an interface to validate the names of drums found in tablatures
    and create the specific representation for each tablature line that can
    play, in some way, those drums.

type DrumTab []TabLine
    A DrumTab represents a sequenced collection of tablature lines.

func ReadTab(rdr io.Reader, kit DrumKit) (DrumTab, error)
    ReadTab reads a tablature from rdr and using a DrumKit, creates a DrumTab
    representing it. If the tablature is in an incorrect format, or refers to
    unsupported or unrecognised drum names, an error will be returned.

    A valid tablature is a text file consisting only of lines in this format: a
    unique two letter drum name, followed by a vertical bar, followed by a
    sequence of pauses and hits. Pauses are represented by '.' or '-', hits are
    represented by any character that is not a pause. An example:-

        HH|..o...o...o...o.
        SD|....o.......o...
        BD|o...o...o...o...

    BUG(jcf): ReadTab only supports the ASCII character set (UTF-8 0..127)

type TabLine interface {
	AddDrum(name string)
	Play(lastLine bool)
}
    TabLine is an interface which can compose (aggregate) the various drums that
    should play at a given point in the tablature. Once all drums are composed,
    the TabLine can be played (ie. the various drums are hit simultaenously).
    When the final TabLine in a DrumTab is played, lastLine will be true.
```
