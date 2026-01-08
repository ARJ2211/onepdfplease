package style

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func RenderTwoFullCols(termWidth, termHeight int, style lipgloss.Style, col1View, col2View string) string {
	spacing := 0
	padding := 0
	height := SplitHeightByPercentage(termHeight, []float64{1}, spacing, padding, 2)
	widths := SplitWidthByPercentage(termWidth, []float64{0.5, 0.5}, spacing, padding, 2)

	// truncateView
	col1View = TruncateView(col1View, widths[0])
	col2View = TruncateView(col2View, widths[1])


	// apply style
	col1View = style.
		Padding(padding).
		Width(widths[0]).
		Height(height[0]).
		Render(col1View)
	col2View = style.
		Padding(padding).
		Width(widths[1]).
		Height(height[0]).
		Render(col2View)

	// add spacing
	spacer := strings.Repeat(" ", spacing)
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		AddSpacerInBetween([]string{col1View, col2View}, spacer)...,
	)
}
