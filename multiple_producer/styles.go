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
)

// color instances
const (
	purpleI  = lipgloss.Color(purple)
	yellowI  = lipgloss.Color(yellow)
	whiteI   = lipgloss.Color(white)
	blueI    = lipgloss.Color(blue)
	redPinkI = lipgloss.Color(redPink)
)

// styles
var (
	mainWin = lipgloss.NewStyle().
		Foreground(whiteI).
		PaddingTop(1).
		PaddingLeft(4).
		Width(120).
		Height(18).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(blueI)

	logWin = lipgloss.NewStyle().
		Foreground(whiteI).
		PaddingLeft(1).
		PaddingRight(1).
		Border(lipgloss.RoundedBorder()).
		AlignVertical(lipgloss.Center).
		BorderForeground(yellowI)
)
