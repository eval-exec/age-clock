package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func getAge() float64 {
	me := Person{BirthDay: time.Date(2000, 0, 0, 0, 0, 0, 0, time.Local)}
	return me.Age()
}

func splitByRow(in []rune) (out [][]rune) {
	var row = 0
	var line []rune
	for _, b := range in {
		if b == '\n' {
			out = append(out, line)
			line = []rune{}
			row++
		} else {
			line = append(line, b)
		}
	}
	return out
}

var X = 242

func ageStr()string {

	return strconv.FormatFloat(getAge(), 'f', 10, 64)
}

func makebox(s tcell.Screen) {
	command := exec.Command("toilet", "-f", "mono12", "-w", "3000", ageStr())
	bs, err := command.CombinedOutput()
	if err != nil {
		panic(err)
	}
	out := splitByRow([]rune(string(bs)))
	lh := len(out)
	ageW := len(out[0])
	w, h := s.Size()
	//
	if w == 0 || h == 0 {
		panic("screen size is invalid ")
		return
	}
	st := tcell.StyleDefault
	if (w - ageW/2) < X {
		X = (w - ageW) / 2
	}

	for row := 0; row < lh; row++ {
		s.SetContent(X, row+(h-lh)/2, out[row][0], out[row][1:], st)
	}
	s.Show()
}

func main() {

	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	s.Clear()

	quit := make(chan struct{})
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(quit)
					return
				case tcell.KeyCtrlL:
					s.Sync()
				}
			case *tcell.EventResize:
				s.Sync()
			}
		}
	}()

	cnt := 0
	dur := time.Duration(0)
loop:
	for {
		select {
		case <-quit:
			break loop
		case <-time.After(time.Millisecond * 50):
		}
		start := time.Now()
		makebox(s)
		cnt++
		dur += time.Now().Sub(start)
	}

	s.Fini()
}
