package main

import (
	"github.com/ungerik/go-progress"
)

func main() {
	bar := progress.Bar{Width: 100, Min: 0, Max: 100, Header: true}
	bar.Newline()
	bar.Draw(50)
	bar.Redraw(25)
	bar.Redraw(75)
	bar.Newline()
}
