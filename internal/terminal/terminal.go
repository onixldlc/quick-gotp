package terminal

import "fmt"

func ClearLine() {
    fmt.Print("\r\033[K")
}

func HideCursor() {
    fmt.Print("\033[?25l")
}

func ShowCursor() {
    fmt.Print("\033[?25h")
}

func MoveCursorDown(lines int) {
    fmt.Printf("\033[%dB", lines)
}

func MoveCursorUp(lines int) {
    fmt.Printf("\033[%dA", lines)
}

func MoveToNextLine() {
    fmt.Print("\033[E")
}

func PositionCursorAtBottom(lineCount int) {
    fmt.Printf("\r\033[%dB\n", lineCount)
    ShowCursor()
}