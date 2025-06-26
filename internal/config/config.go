package config

import (
    "encoding/json"
    "flag"
    "io/ioutil"
    "log"
    _ "embed"
)

//go:embed secret.json.example
var defaultConfig []byte

type Credential struct {
    Secret string `json:"secret"`
    Name   string `json:"name"`
    Delay  int    `json:"delay"`
}

type Config struct {
    Credentials []Credential `json:"credentials"`
    OneTime     bool
}

func LoadConfig() Config {
    tempFlag := flag.Bool("temp", false, "Use temporary mode with a provided secret")
    secretArg := flag.String("secret", "", "The secret to use in temporary mode")
    delayArg := flag.Int("delay", 30, "Time period in seconds for OTP refresh (default is 30 seconds)")
    oneTimeFlag := flag.Bool("one-time", false, "Use one-time mode (not implemented yet)")
    flag.Parse()

    if *oneTimeFlag && !*tempFlag {
        log.Fatal("One-time mode (--one-time) can only be used with temporary mode (--temp)")
    }

    if *tempFlag {
        return loadTempConfig(*secretArg, *delayArg, *oneTimeFlag)
    }
    return loadFileConfig()
}

func loadTempConfig(secret string, delay int, oneTime bool) Config {
    if secret == "" {
        log.Fatal("Secret must be provided in temporary mode")
    }
    
    return Config{
        Credentials: []Credential{
            {
                Secret: secret,
                Name:   "Temporary OTP",
                Delay:  delay,
            },
        },
        OneTime: oneTime,
    }
}

func loadFileConfig() Config {
    var cfg Config
    
    file, err := ioutil.ReadFile("secret.json")
    if err != nil {
        log.Fatalf("Error reading secret.json: %v", err)

		// create default secret.json if it doesn't exist
		if err := ioutil.WriteFile("secret.json", defaultConfig, 0644); err != nil {
			log.Fatalf("Error creating secret.json: %v", err)
		}
	}

    if err := json.Unmarshal(file, &cfg); err != nil {
        log.Fatalf("Error parsing secret.json: %v", err)
    }
    
    return cfg
}