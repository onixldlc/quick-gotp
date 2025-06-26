# Quick-GOTP

Quick-GOTP is a simple Go application that generates One-Time Passwords (OTPs) using the TOTP algorithm. The application can read configuration settings from a JSON file or accept temporary settings via command-line arguments.

## Features

- Generates OTPs at a refresh rate of 4 frames per second.
- Supports reading from a `secret.json` file containing the fields:
  - `secret`: The shared secret used for OTP generation.
  - `name`: A descriptive name for the OTP.
  - `delay`: The time interval (in seconds) for OTP refresh (default is 30 seconds).
- Temporary OTP generation using the `--temp` flag, which allows users to specify a secret directly from the command line.

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/quick-gotp.git
   cd quick-gotp
   ```

2. Install the required dependencies:
   ```
   go mod tidy
   ```

## Configuration

Create a `secret.json` file in the root directory with the following structure:

```json
{
  "secret": "YOUR_SECRET_HERE",
  "name": "YOUR_NAME_HERE",
  "delay": 30
}
```

## Usage

To run the application normally, use:
```
go run cmd/main.go
```

To run the application with a temporary secret, use:
```
go run cmd/main.go --temp YOUR_TEMP_SECRET --delay YOUR_DELAY_IN_SECONDS
```

If the `--delay` flag is not provided, it defaults to 30 seconds.

## License

This project is licensed under the MIT License. See the LICENSE file for details.