package sequencer

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
	"time"
)

var (
	egTab1 = "D1|\n"        // invalid format
	egTab2 = "D3|.\n"       // unknown drum
	egTab3 = "D1|.\nD1|o\n" // duplicated drum
	// ReadTab should produce this...
	egTab4 = "D1|..o.\n" + "D2|o.o.\n"
	okTab4 = DrumTab{
		&testLine{"D2", nil},
		&testLine{"", nil},
		&testLine{"D1+D2", nil},
		&testLine{"", nil},
	}
	// PlayTab should produce this...
	egTab5 = "D1|.o.o.o.o.o.o.o.o\n" + "D2|.o.o.o.o.o.o.o.o\n"
	okTab5 = "_D1+D2_D1+D2_D1+D2_D1+D2_D1+D2_D1+D2_D1+D2_D1+D2\n"
)

func TestReadTab(t *testing.T) {
	kit := testKit{nil}
	rdr := bytes.NewReader([]byte(egTab1))
	if _, err := ReadTab(rdr, kit); !errors.Is(err, ErrInvalidTab) {
		t.Errorf("got %v, expected %v", err, ErrInvalidTab)
	}
	rdr = bytes.NewReader([]byte(egTab2))
	if _, err := ReadTab(rdr, kit); !errors.Is(err, ErrUnknownDrum) {
		t.Errorf("got %v, expected %v", err, ErrUnknownDrum)
	}
	rdr = bytes.NewReader([]byte(egTab3))
	if _, err := ReadTab(rdr, kit); !errors.Is(err, ErrInvalidTab) {
		t.Errorf("got %v, expected %v", err, ErrInvalidTab)
	}
	rdr = bytes.NewReader([]byte(egTab4))
	tab, err := ReadTab(rdr, kit)
	if err != nil {
		t.Errorf("got %v, expected nil", err)
	}
	if !reflect.DeepEqual(tab, okTab4) {
		t.Errorf("got %+v, expected %+v", tab, okTab4)
	}
}

func TestPlayTab(t *testing.T) {
	var out bytes.Buffer
	kit := testKit{&out}
	rdr := bytes.NewReader([]byte(egTab5))
	tab, err := ReadTab(rdr, kit)
	if err != nil {
		t.Fatalf("got %v, expected nil", err)
	}
	// The tempo is wavering at best, so just ensure it's not more than a
	// half-second slow or fast.
	now := time.Now()
	PlayTab(tab, 60, false)
	dur := time.Since(now)
	if dur <= 7*time.Second/2 || dur >= 9*time.Second/2 {
		t.Errorf("got %v, expected %v", dur, LinesPerBeat*time.Second)
	}
	if bytes.Compare(out.Bytes(), []byte(okTab5)) != 0 {
		t.Errorf("got %+v, expected %+v", string(out.Bytes()), okTab5)
	}
}

type testKit struct {
	w io.Writer
}

func (tk testKit) HasDrum(name string) bool {
	return name == "D1" || name == "D2"
}

func (tk testKit) NewLine() TabLine {
	return &testLine{"", tk.w}
}

type testLine struct {
	val string
	w   io.Writer
}

func (tl *testLine) AddDrum(name string) {
	if tl.val == "" {
		tl.val = name
		return
	}
	// Must overcome ReadTab's indeterminate map key iteration, to
	// reliably compare the test's output against what is expected.
	if strings.Compare(tl.val, name) < 0 {
		tl.val += "+" + name
	} else {
		tl.val = name + "+" + tl.val
	}
}

func (tl *testLine) Play(lastLine bool) {
	var s string
	if tl.val == "" {
		s = "_"
	} else {
		s = tl.val
	}
	if _, err := tl.w.Write([]byte(s)); err != nil {
		panic(err)
	}
	if lastLine {
		if _, err := tl.w.Write([]byte{'\n'}); err != nil {
			panic(err)
		}
	}
}

// This helps make test failures more readable.
func (tl *testLine) String() string {
	return tl.val
}
