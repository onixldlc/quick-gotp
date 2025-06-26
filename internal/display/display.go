package display

import (
    "fmt"
    "time"

    "github.com/quick-gotp/v2/internal/config"
    "github.com/quick-gotp/v2/internal/otp"
    "github.com/quick-gotp/v2/internal/terminal"
)

type OTPDisplay struct {
    config      config.Config
    otpMap      map[string]string
    lastRefresh map[string]time.Time
    neededLines int
}

func New(cfg config.Config) *OTPDisplay {
    return &OTPDisplay{
        config:      cfg,
        otpMap:      make(map[string]string),
        lastRefresh: make(map[string]time.Time),
        neededLines: 1 + (len(cfg.Credentials) * 4), // header + (3 lines per cred + blank line)
    }
}

func (d *OTPDisplay) PrepareTerminal() {
    terminal.HideCursor()
    
    for i := 0; i < d.neededLines; i++ {
        fmt.Println()
    }
    
    terminal.MoveCursorUp(d.neededLines)
    
    terminal.ClearLine()
    fmt.Println("otp-codes:")
}

func (d *OTPDisplay) GetNeededLines() int {
    return d.neededLines
}

// Update refreshes the OTP display
func (d *OTPDisplay) Update() {
    now := time.Now()
    
    // Move to the line after the header
    fmt.Print("\033[1B")
    
    for i, cred := range d.config.Credentials {
        if otp.ShouldRefresh(d.lastRefresh[cred.Name], cred.Delay) {
            d.otpMap[cred.Name] = otp.Generate(cred.Secret)
            d.lastRefresh[cred.Name] = now
        }
        
        remaining := otp.TimeRemaining(cred.Delay)
        
        terminal.ClearLine()
        fmt.Printf("  - name: %s", cred.Name)
        
        terminal.MoveToNextLine()
        terminal.ClearLine()
        fmt.Printf("      - code: %s", d.otpMap[cred.Name])
        
        terminal.MoveToNextLine()
        terminal.ClearLine()
        fmt.Printf("      - timeleft: %d", remaining)
        
        if i < len(d.config.Credentials) - 1 {
            terminal.MoveToNextLine()
            terminal.ClearLine()
            fmt.Print("")
        }
        
        terminal.MoveToNextLine()
    }
    
    terminal.MoveCursorUp(1 + (len(d.config.Credentials) * 4) - 1)
}

func DisplayOneTime(cred config.Credential) {
    otpCode := otp.Generate(cred.Secret)
    timeLeft := otp.TimeRemaining(cred.Delay)
    fmt.Printf("%s:%d\n", otpCode, timeLeft)
}