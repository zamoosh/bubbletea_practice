package main

import (
	"github.com/charmbracelet/lipgloss"
)

// color codes
const (
	purple  = "#3a0ca3"
	yellow  = "#ffc300"
	blue    = "#0079FF"
	white   = "#ffffff"
	redPink = "#c21362"
	green   = "#06d6a0"
)

// color instances
const (
	purpleI  = lipgloss.Color(purple)
	yellowI  = lipgloss.Color(yellow)
	whiteI   = lipgloss.Color(white)
	blueI    = lipgloss.Color(blue)
	redPinkI = lipgloss.Color(redPink)
	greenI   = lipgloss.Color(green)
)

// styles
var (
	mainWin = lipgloss.NewStyle().
		Foreground(whiteI).
		PaddingTop(1).
		PaddingLeft(4).
		PaddingRight(4).
		Width(120).
		Height(20).
		Border(lipgloss.RoundedBorder())

	logWin = lipgloss.NewStyle().
		Foreground(whiteI).
		Height(15).
		PaddingLeft(1).
		PaddingRight(1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(yellowI)

	imeiListStyle = lipgloss.NewStyle().
			Foreground(greenI).
			Bold(true)
)
