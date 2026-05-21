package display

import (
	"fmt"
	"strings"
	"time"
	"github.com/charmbracelet/lipgloss"
	"github.com/plehmann/fl-launches/api"
)

const wrapWidth = 64

var (
	orange    = lipgloss.Color("#FF6B00")
	white     = lipgloss.Color("#FFFFFF")
	dimGray   = lipgloss.Color("#888888")
	accent    = lipgloss.Color("#00D4FF")
	green     = lipgloss.Color("#00FF9F")
	yellow    = lipgloss.Color("#FFD700")

	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(white).
			Background(orange).
			Padding(0, 2)

	numberStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(orange)

	missionNameStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(white)

	vehicleStyle = lipgloss.NewStyle().
			Foreground(accent)

	padStyle = lipgloss.NewStyle().
			Foreground(dimGray)

	dateStyle = lipgloss.NewStyle().
			Foreground(yellow)

	countdownStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(green)

	descStyle = lipgloss.NewStyle().
			Foreground(dimGray)

	dividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#333333"))
)

func Render(launches []api.Launch) {
	fmt.Println()
	fmt.Println(headerStyle.Render("  🚀  CAPE CANAVERAL UPCOMING LAUNCHES  "))
	fmt.Println()

	if len(launches) == 0 {
		fmt.Println("  No upcoming launches found.")
		fmt.Println()
		return
	}

	for i, l := range launches {
		renderCard(i+1, l)
		if i < len(launches)-1 {
			fmt.Println(dividerStyle.Render("  " + strings.Repeat("─", wrapWidth)))
		}
	}
	fmt.Println()
	fmt.Println(padStyle.Render("  Data: Launch Library 2 (thespacedevs.com)"))
	fmt.Println()
}

func renderCard(n int, l api.Launch) {
	now := time.Now()
	countdown := formatCountdown(l.Net, now)
	localTime := l.Net.Local().Format("Mon Jan 2, 2006 at 3:04 PM MST")

	fmt.Printf("  %s  %s\n",
		numberStyle.Render(fmt.Sprintf("#%d", n)),
		missionNameStyle.Render(strings.ToUpper(l.Name)),
	)

	vehicle := l.Rocket.Configuration.FullName
	if vehicle == "" {
		vehicle = "Unknown vehicle"
	}
	padName := l.Pad.Name
	if padName == "" {
		padName = l.Pad.Location.Name
	}
	fmt.Printf("      %s  %s  %s\n",
		vehicleStyle.Render(vehicle),
		padStyle.Render("·"),
		padStyle.Render(padName),
	)

	fmt.Printf("      %s  %s  %s\n",
		dateStyle.Render(localTime),
		padStyle.Render("|"),
		countdownStyle.Render(countdown),
	)

	desc := "No mission description available."
	if l.Mission != nil && l.Mission.Description != "" {
		desc = l.Mission.Description
	}
	fmt.Println()
	for _, line := range strings.Split(wordWrap(desc, wrapWidth-6), "\n") {
		fmt.Println("      " + descStyle.Render(line))
	}
	fmt.Println()
}

func formatCountdown(net time.Time, now time.Time) string {
	diff := net.Sub(now)
	prefix := "T-"
	if diff < 0 {
		prefix = "T+"
		diff = -diff
	}

	days := int(diff.Hours()) / 24
	hours := int(diff.Hours()) % 24
	mins := int(diff.Minutes()) % 60

	if days > 0 {
		return fmt.Sprintf("%s %dd %dh %dm", prefix, days, hours, mins)
	}
	if hours > 0 {
		return fmt.Sprintf("%s %dh %dm", prefix, hours, mins)
	}
	return fmt.Sprintf("%s %dm", prefix, mins)
}

func wordWrap(s string, width int) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return s
	}

	var lines []string
	currentLine := words[0]

	for _, word := range words[1:] {
		if len([]rune(currentLine))+1+len([]rune(word)) <= width {
			currentLine += " " + word
		} else {
			lines = append(lines, currentLine)
			currentLine = word
		}
	}
	lines = append(lines, currentLine)
	return strings.Join(lines, "\n")
}


