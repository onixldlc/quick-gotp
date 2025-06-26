package otp

import (
    "log"
    "time"

    "github.com/pquerna/otp/totp"
)

func Generate(secret string) string {
    code, err := totp.GenerateCode(secret, time.Now())
    if err != nil {
        log.Fatalf("Error generating OTP: %v", err)
    }
    return code
}

func TimeRemaining(period int) int {
    now := time.Now().Unix()
    remaining := period - int(now%int64(period))
    if remaining < 0 {
        return 0
    }
    return remaining
}

func ShouldRefresh(lastRefresh time.Time, period int) bool {
    if lastRefresh.IsZero() {
        return true
    }
    
    now := time.Now()
    return now.Unix()/int64(period) != lastRefresh.Unix()/int64(period)
}