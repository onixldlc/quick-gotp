package main

import (
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/quick-gotp/v2/internal/config"
    "github.com/quick-gotp/v2/internal/display"
    "github.com/quick-gotp/v2/internal/terminal"
)

func main() {
    cfg := config.LoadConfig()

    // Handle one-time mode
    if cfg.OneTime {
        display.DisplayOneTime(cfg.Credentials[0])
        return
    }

    // Hide the cursor and ensure it's restored on exit
    terminal.HideCursor()
    defer terminal.ShowCursor()
    
    otpDisplay := display.New(cfg)
    
    // Set up signal handling to restore cursor on interrupt
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        terminal.PositionCursorAtBottom(otpDisplay.GetNeededLines())
        os.Exit(0)
    }()
    
    otpDisplay.PrepareTerminal()
    
    refreshRate := 82 * time.Millisecond
    ticker := time.NewTicker(refreshRate)
    defer ticker.Stop()
    
    // Main loop
    for {
        select {
        case <-ticker.C:
            otpDisplay.Update()
        }
    }
}