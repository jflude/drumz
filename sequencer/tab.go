// Package sequencer provides interfaces and functions to read and play
// drum tablatures.
package sequencer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"time"
)

// LinesPerBeat determines how many lines in a tablature correspond to a
// single beat.
const LinesPerBeat = 4

// DrumKit is an interface to validate the names of drums found in tablatures
// and create the specific representation for each tablature line that can
// play, in some way, those drums.
type DrumKit interface {
	HasDrum(name string) bool
	NewLine() TabLine
}

// A DrumTab represents a sequenced collection of tablature lines.
type DrumTab []TabLine

// TabLine is an interface which can compose (aggregate) the various drums
// that should play at a given point in the tablature.  Once all drums are
// composed, the TabLine can be played (ie. the various drums are hit
// simultaenously).  When the final TabLine in a DrumTab is played, lastLine
// will be true.
type TabLine interface {
	AddDrum(name string)
	Play(lastLine bool)
}

// ErrInvalidTab is returned by ReadTab for tablatures with incorrect formats.
var ErrInvalidTab = errors.New("invalid tablature")

// ErrUnknownDrum is returned by ReadTab for tablatures with drums which
// DrumKit does not support or recognise.
var ErrUnknownDrum = errors.New("unknown drum")

// PlayTab takes a DrumTab created by ReadTab and plays it, in some way, at
// the tempo given by the bpm argument.  If repeat is false it will be played
// just once.
func PlayTab(tab DrumTab, bpm int, repeat bool) {
	period := time.Minute / time.Duration(bpm*LinesPerBeat)
	for {
		for i := 0; i < len(tab); i++ {
			now := time.Now()
			tab[i].Play(i == len(tab)-1)
			// for me, Sleep() gives steadier tempos than Tick()
			time.Sleep(time.Until(now.Add(period)))
		}
		if !repeat {
			break
		}
	}
}

// ReadTab reads a tablature from rdr and using a DrumKit, creates a DrumTab
// representing it.  If the tablature is in an incorrect format, or refers
// to unsupported or unrecognised drum names, an error will be returned.
//
// A valid tablature is a text file consisting only of lines in this format:
// a unique two letter drum name, followed by a vertical bar, followed by a
// sequence of pauses and hits.  Pauses are represented by '.' or '-', hits
// are represented by any character that is not a pause.  An example:-
//
//         HH|..o...o...o...o.
//         SD|....o.......o...
//         BD|o...o...o...o...
//
// BUG(jcf): ReadTab only supports the ASCII character set (UTF-8 0..127)
func ReadTab(rdr io.Reader, kit DrumKit) (DrumTab, error) {
	m := make(map[string]string)
	width := 0
	lineNo := 0
	sc := bufio.NewScanner(rdr)
	for sc.Scan() {
		lineNo++
		txt := sc.Text()
		// Is line prefix valid?
		if len(txt) < 4 || txt[2] != '|' {
			return nil, fmt.Errorf("error: %w: line %v",
				ErrInvalidTab, lineNo)
		}
		// Is the drum name valid?
		name, pat := txt[:2], txt[3:]
		if !kit.HasDrum(name) {
			return nil, fmt.Errorf("error: %w: %q",
				ErrUnknownDrum, name)
		}
		// Is the drum name unique for this tablature?
		if _, exists := m[name]; exists {
			return nil, fmt.Errorf("error: %w: line %v",
				ErrInvalidTab, lineNo)
		}
		m[name] = pat
		if len(pat) > width {
			width = len(pat)
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	// Does the tablature actually contain any lines?
	if len(m) == 0 || width == 0 {
		return nil, ErrInvalidTab
	}
	// For each line, compose all the drums hit on that line.
	tab := makeTab(kit, width)
	for name, pat := range m {
		for i := 0; i < width; i++ {
			if i < len(pat) && pat[i] != '.' && pat[i] != '-' {
				tab[i].AddDrum(name)
			}
		}
	}
	return tab, nil
}

func makeTab(kit DrumKit, width int) DrumTab {
	// X may implement Y but that does not mean []X implements []Y
	t := make([]TabLine, width)
	for i, _ := range t {
		t[i] = kit.NewLine()
	}
	return t
}
