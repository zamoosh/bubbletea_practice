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
	cyan    = "#0FFFFF"
)

// color instances
const (
	purpleI  = lipgloss.Color(purple)
	yellowI  = lipgloss.Color(yellow)
	whiteI   = lipgloss.Color(white)
	blueI    = lipgloss.Color(blue)
	redPinkI = lipgloss.Color(redPink)
	cyanI    = lipgloss.Color(cyan)
)

const (
	infoSymbol = "ℹ"
	okSymbol = "✔"
	warnSymbol = "⚠"
	errorSymbol = "✖"
)

// styles
var (
	mainWin = lipgloss.NewStyle().
		Foreground(whiteI).
		PaddingTop(1).
		PaddingLeft(4).
		PaddingRight(4).
		Width(150).
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
			Foreground(cyanI).
			Bold(true)

	infoStyle = lipgloss.NewStyle().
			Background(purpleI).
			Foreground(whiteI).Render

	dangerStyle = lipgloss.NewStyle().
		Background(redPinkI).
		Foreground(whiteI).Render
)
