This Go module provides a simple, customizable progress bar for your console applications. It is easy to use and lightweight.

## Features

- Customizable bar width
- Displays elapsed time
- Incremental updates

## Usage

First, import the module:

```go
import "github.com/raulfilimon/golang-progressbar.git"
```

Create a new progress bar with the total number of items and the desired bar width:

```go
bar := progressbar.NewProgressBar(100, 50)
```

Increment the progress bar as tasks are completed:

```go
bar.Increment()
```

The progress bar is automatically rendered after each increment.

## Classes and Methods

### ProgressBar

The `ProgressBar` struct represents a progress bar. It has the following fields:

- `Total`: The total number of items.
- `Current`: The current number of items.
- `BarWidth`: The width of the progress bar.
- `StartTime`: The time when the progress bar was created.

### NewProgressBar(total int, barWidth int) *ProgressBar

This function creates a new progress bar. It takes the total number of items and the desired bar width as arguments.

### Increment()

This method increases the current number of items by one and renders the progress bar.

### Render()

This method renders the progress bar. It calculates the progress, creates the bar string, calculates the elapsed time, and prints the bar, current items, total items, and elapsed time.

### stringOfChar(n int, char string) string

This helper function returns a string of a given character repeated n times. It is used to create the bar string.

## Note

This module uses the `fmt` and `time` standard libraries.