package format

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/term"
)

const (
	defaultTerminalWidth = 80
	progressBarScale     = 0.5 // Bar width is 50% of terminal width
)

// ANSI Catppuccin Mocha palette
type color string

const (
	// Special
	RESET = "\033[0m"
	BOLD  = "\033[1m"
	// Base colors
	ROSEWATER color = "\033[38;2;245;224;220m"
	FLAMINGO  color = "\033[38;2;242;205;205m"
	PINK      color = "\033[38;2;245;194;231m"
	MAUVE     color = "\033[38;2;203;166;247m"
	RED       color = "\033[38;2;243;139;168m"
	MAROON    color = "\033[38;2;235;160;172m"
	PEACH     color = "\033[38;2;250;179;135m"
	YELLOW    color = "\033[38;2;249;226;175m"
	GREEN     color = "\033[38;2;166;227;161m"
	TEAL      color = "\033[38;2;148;226;213m"
	SKY       color = "\033[38;2;137;220;235m"
	SAPPHIRE  color = "\033[38;2;116;199;236m"
	BLUE      color = "\033[38;2;137;180;250m"
	LAVENDER  color = "\033[38;2;180;190;254m"
	TEXT      color = "\033[38;2;205;214;244m"
	SUBTEXT1  color = "\033[38;2;186;194;222m"
	SUBTEXT0  color = "\033[38;2;166;173;200m"
	OVERLAY2  color = "\033[38;2;147;153;178m"
	OVERLAY1  color = "\033[38;2;127;132;156m"
	OVERLAY0  color = "\033[38;2;108;112;134m"
	SURFACE2  color = "\033[38;2;88;91;112m"
	SURFACE1  color = "\033[38;2;69;71;90m"
	SURFACE0  color = "\033[38;2;49;50;68m"
	BASE      color = "\033[38;2;30;30;46m"
	MANTLE    color = "\033[38;2;24;24;37m"
	CRUST     color = "\033[38;2;17;17;27m"
)

var ansiPattern = regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)

// getTerminalWidth returns the current terminal width or falls back to default
func GetTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("error: getting term width: progressing with default: %w", err)
		return defaultTerminalWidth
	}

	return width
}

// alignCenter centers the given text within the specified width
func AlignCenter(text string, width int) string {
	// stripping the ansi chars for padding calculations
	strippedText := ansiPattern.ReplaceAllString(text, "")
	textLen := len(strippedText)
	if textLen >= width {
		return strippedText
	}
	padding := (width - textLen) / 2

	// add the fullText back in
	return strings.Repeat(" ", padding) + text
}

func Label(label string, value string, maxLabelLen, maxValueLen int, colors ...color) string {
	// Format the value with proper padding
	paddedValue := strings.Repeat(" ", maxValueLen-len(value)) + value

	// Format the label with proper padding
	paddedLabel := label + strings.Repeat(" ", maxLabelLen-len(label))

	if len(colors) == 1 {
		paddedLabel = string(colors[0]) + paddedLabel + string(RESET)
		paddedValue = string(colors[0]) + paddedValue + string(RESET)
	}
	if len(colors) == 2 {
		paddedLabel = string(colors[0]) + paddedLabel + string(RESET)
		paddedValue = string(colors[1]) + paddedValue + string(RESET)
	}

	return fmt.Sprintf("%s %s", paddedLabel, paddedValue)
}

func Text(labelConfig map[string]string) []string {
	termWidth := GetTerminalWidth()

	// Calculate padding for keys and values
	maxLabelLen := 0
	maxValueLen := 0
	for label, value := range labelConfig {
		if len(label) > maxLabelLen {
			maxLabelLen = len(label)
		}
		if len(value) > maxValueLen {
			maxValueLen = len(value)
		}
	}

	var statusLines []string
	for label, value := range labelConfig {
		statusLines = append(statusLines, AlignCenter(Label(label, value, maxLabelLen, maxValueLen, LAVENDER, LAVENDER), termWidth))
	}

	return statusLines
}

// createProgressBar generates a progress bar string
// bar alterative icons
// ██░ | ▣ ⬚ | ● ○ ○ | [= ]
func CreateProgressBar(current, total int) string {
	width := int(float64(GetTerminalWidth()) * progressBarScale)
	completed := int(float64(current) / float64(total) * float64(width))
	bar := strings.Builder{}
	bar.WriteString("[")
	bar.WriteString(strings.Repeat("▣ ", completed))
	bar.WriteString(strings.Repeat("⬚ ", width-completed))
	bar.WriteString("]")
	return bar.String()
}

// TODO: if term supports font icons else use [=]
func CreateProgressBarWithColor(current, total int, barColor color) string {
	termWidth := GetTerminalWidth()
	barWidth := int(float64(termWidth) * progressBarScale)
	completed := int(float64(current) / float64(total) * float64(barWidth-2)) // -2 for brackets

	bar := strings.Builder{}
	bar.WriteString(strings.Repeat("=", completed))
	bar.WriteString(strings.Repeat(" ", barWidth-2-completed))
	// bar.WriteString(strings.Repeat("█", completed))
	// bar.WriteString(strings.Repeat("░", barWidth-2-completed))

	// Calculate padding for centering
	totalBarWidth := barWidth // includes brackets
	padding := (termWidth - totalBarWidth) / 2

	// return fmt.Sprintf("%s%s%s%s",
	// 	strings.Repeat(" ", padding),
	// 	barColor,
	// 	bar.String(),
	// 	RESET)

	return fmt.Sprintf("%s%s[%s]%s",
		strings.Repeat(" ", padding),
		barColor,
		bar.String(),
		RESET)
}
