package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var style = map[string]lipgloss.Style{
	"RB": lipgloss.NewStyle().
		Background(lipgloss.Color("#D3D3D3")). // light gray
		Foreground(lipgloss.Color("#000000")).
		Padding(0, 2),
	"Bus": lipgloss.NewStyle().
		Background(lipgloss.Color("#800080")). // purple
		Foreground(lipgloss.Color("#ffffff")).
		Padding(0, 2),
	"S": lipgloss.NewStyle().
		Background(lipgloss.Color("#008004")). // green
		Foreground(lipgloss.Color("#ffffff")).
		Padding(0, 2),
	"Walk": lipgloss.NewStyle().
		Background(lipgloss.Color("#004980")). // green
		Foreground(lipgloss.Color("#ffffff")).
		Padding(0, 2),
	"Border": lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#3C3C3C")).Width(50).PaddingTop(1).
		PaddingRight(4).PaddingBottom(1).PaddingLeft(4),
}

func (m model) TripView() string {
	view := []string{}
	for _, trip := range m.data {
		parts := []string{}
		header := ""

		for _, part := range trip.Abschnitte {
			newPart := formatAbschnitt(part.Dauer*30/trip.Dauer, part.Art, part.Text)
			parts = append(parts, newPart)
		}
		if len(parts) > 3 {
			top_parts := lipgloss.JoinHorizontal(lipgloss.Top, parts[:3]...)
			bottom_parts := lipgloss.JoinHorizontal(lipgloss.Top, parts[3:]...)
			header = lipgloss.JoinVertical(lipgloss.Top, top_parts, bottom_parts)
		} else {
			header = lipgloss.JoinHorizontal(lipgloss.Top, parts...)
		}

		duration := lipgloss.NewStyle().Foreground(lipgloss.Color("5")).Render(fmt.Sprintf("󰞌 %02d:%02d", trip.Dauer/60, trip.Dauer%60))
		label := fmt.Sprintf("%v ➝ %v", trip.Abschnitte[0].AbfahrtsZeitpunkt.Format("15:04"), trip.Abschnitte[len(trip.Abschnitte)-1].AnkunftsZeitpunkt.Format("15:04"))

		view = append(view, style["Border"].Render(lipgloss.JoinVertical(lipgloss.Left, header, duration, label)))
	}

	return lipgloss.JoinVertical(lipgloss.Left, view...)
}

func formatAbschnitt(perc int, art string, text string) string {
	if len(text) > perc {
		perc += 3
	}
	abschnittStyle, ok := style[art]
	if !ok {
		if text == "" {
			abschnittStyle = style["Walk"]
			text = ""
		} else {
			abschnittStyle = style["RB"]
		}
	}

	text = strings.Replace(text, "Bus", "", 1)
	return abschnittStyle.Width(perc).MarginRight(1).Render(text)
}
