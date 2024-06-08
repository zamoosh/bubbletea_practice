package main

import (
	"github.com/charmbracelet/lipgloss"
)

// color codes
const (
	purple  = "#8338ec"
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
		Height(23).
		Border(lipgloss.RoundedBorder())

	logWin = lipgloss.NewStyle().
		Foreground(whiteI).
		Height(20).
		PaddingLeft(1).
		PaddingRight(1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(yellowI)

	imeiListStyle = lipgloss.NewStyle().
			Foreground(greenI).
			Bold(true)

	infoStyle = lipgloss.NewStyle().
			Background(purpleI).
			Foreground(whiteI).Render

	dangerStyle = lipgloss.NewStyle().
		Background(redPinkI).
		Foreground(whiteI).Render
)
