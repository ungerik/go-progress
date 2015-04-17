package progress

import (
	"fmt"
)

type Bar struct {
	Width  int
	Min    float32
	Max    float32
	Divide int
	Marker rune
	Bar    rune
	Header bool
}

func (bar *Bar) Draw(pos float32) *Bar {
	width := bar.Width
	if width == 0 {
		width = 50
	}
	marker := bar.Marker
	if marker == 0 {
		marker = '|'
	}
	line := bar.Bar
	if line == 0 {
		line = '='
	}

	if bar.Header {
		div := bar.Divide
		if div == 0 {
			div = 4
		}
		t := bar.Min
		for i := 0; i < width; i++ {
			if i == int(t-bar.Min+0.5) || i == width-1 {
				fmt.Print(string(marker))
				t += (bar.Max - bar.Min) / float32(div)
			} else {
				fmt.Print("-")
			}
		}
		bar.Newline()
	}

	p := (pos-bar.Min)/(bar.Max-bar.Min)*float32(bar.Width) + 0.5

	for i := 0; i < width; i++ {
		if i == 0 || i == width-1 {
			fmt.Print(string(marker))
		} else if i <= int(p) {
			fmt.Print(string(line))
		} else {
			fmt.Print(" ")
		}
	}

	return bar
}

func (bar *Bar) Redraw(pos float32) *Bar {
	width := bar.Width
	if width == 0 {
		width = 50
	}
	for i := 0; i < width; i++ {
		fmt.Print("\b")
	}
	if bar.Header {
		bar.UndoNewline()
	}
	return bar.Draw(pos)
}

func (bar *Bar) Newline() *Bar {
	fmt.Print("\n")
	return bar
}

func (bar *Bar) UndoNewline() *Bar {
	fmt.Print("\033[F")
	return bar
}
