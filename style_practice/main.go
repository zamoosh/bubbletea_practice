package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	text := "ali ali ali"

	st := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00f5d4")).
		PaddingTop(5).
		PaddingBottom(5).
		PaddingLeft(10).
		Width(50).
		Border(lipgloss.ThickBorder())

	st2 := lipgloss.NewStyle().
		// Foreground(lipgloss.Color("#FFFFFF")).
		PaddingTop(1).
		PaddingBottom(1).
		PaddingLeft(2).
		PaddingRight(3).
		// Width(20).
		Border(lipgloss.RoundedBorder())

	fmt.Println(st.Render(st2.Render(text)))
	fmt.Println(st.Render(text))
}
