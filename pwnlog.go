package pwnlog

import (
    "os"
    "github.com/fatih/color"
)

// Define the logging levels as an enum
const (
    DebugLevel int = iota
    InfoLevel
    WarningLevel
    ErrorLevel
    FatalLevel
)


// Log levels:
//   debug    -> [*]
//   info     -> [+]
//   warning  -> [-]
//   error    -> [!]
//   fatal    -> [!!] (followed by death of program)

var prefix = []string{
    DebugLevel:    "[*] ",
    InfoLevel:     "[+] ",
    WarningLevel:  "[-] ",
    ErrorLevel:    "[!] ",
    FatalLevel:    "[!!] ",
}

// Logger struct
type Logger struct {
    logLevel int
    logColors []*color.Color
}

// Create a new logger
func New(level int) *Logger {
    // Check for sane level
    if level < DebugLevel || level > FatalLevel {
        // garbage in, you get debug level
        level = DebugLevel
    }

    // create colors for logging levels
    colors := []*color.Color{}

    debugColor := color.New(color.FgGreen)
    colors = append(colors, debugColor)
    
    infoColor := color.New(color.FgWhite)
    colors = append(colors, infoColor)

    warningColor := color.New(color.FgYellow)
    colors = append(colors, warningColor)
    
    errorColor := color.New(color.FgRed)
    colors = append(colors, errorColor)
    
    fatalColor := color.New(color.FgRed, color.Bold)
    colors = append(colors, fatalColor)
    
    return &Logger{level, colors}
}

func (l *Logger) Debug(format string, v ...interface{}) {
    if l.logLevel != DebugLevel {
        return
    }
    l.logColors[DebugLevel].Printf(prefix[DebugLevel]+format+"\n", v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
    if l.logLevel > InfoLevel {
        return
    }
    l.logColors[InfoLevel].Printf(prefix[InfoLevel]+format+"\n", v...)
}

func (l *Logger) Warning(format string, v ...interface{}) {
    if l.logLevel > WarningLevel {
        return
    }
    l.logColors[WarningLevel].Printf(prefix[WarningLevel]+format+"\n", v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
    if l.logLevel > ErrorLevel {
        return
    }
    l.logColors[ErrorLevel].Printf(prefix[ErrorLevel]+format+"\n", v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
    if l.logLevel > FatalLevel {
        return
    }
    l.logColors[FatalLevel].Printf(prefix[FatalLevel]+format+"\n", v...)
    os.Exit(-1)
}

func (l *Logger) SetLevel(level int) {
    // Check for sane level
    if level < DebugLevel && level > FatalLevel {
        // garbage in, you get debug level
        level = DebugLevel
    }
    l.logLevel = level
}
