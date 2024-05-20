package golangprogressbar

import (
	"fmt"
	"time"
)

// ProgressBar represents a progress bar
type ProgressBar struct {
	Total     int       // Total is the total number of items
	Current   int       // Current is the current number of items
	BarWidth  int       // BarWidth is the width of the progress bar
	StartTime time.Time // StartTime is the time when the progress bar was created
}

// NewProgressBar creates a new progress bar
func NewProgressBar(total int, barWidth int) *ProgressBar {
	return &ProgressBar{
		Total:     total,
		Current:   0,
		BarWidth:  barWidth,
		StartTime: time.Now(),
	}
}

// Increment increases the current number of items
func (p *ProgressBar) Increment() {
	p.Current++
	p.Render()
}

// Render renders the progress bar
func (p *ProgressBar) Render() {
	progress := (p.Current * p.BarWidth) / p.Total
	bar := "[" + stringOfChar(progress, "=") + stringOfChar(p.BarWidth-progress, " ") + "]"
	elapsed := time.Since(p.StartTime)
	fmt.Printf("\r%s %d/%d Time elpased: %s", bar, p.Current, p.Total, elapsed.String())
}

// stringOfChar returns a string of a given character repeated n times
func stringOfChar(n int, char string) string {
	var result string
	for i := 0; i < n; i++ {
		result += char
	}
	return result
}
