package logger

import (
	"encoding/json"
	"os"
	"runtime"
	"sync"
	"time"
)

type brush func(string) string

func newBrush(color string) brush {
	pre := "\033["
	reset := "\033[0m"
	return func(text string) string {
		return pre + color + "m" + text + reset
	}
}

var colors = []brush{
	newBrush("1;41"), // Emergency          red
	newBrush("1;35"), // Alert              purple
	newBrush("1;34"), // Critical           blue
	newBrush("1;31"), // Error              red
	newBrush("1;33"), // Warn               yellow
	newBrush("1;36"), // Informational      sky blue
	newBrush("1;32"), // Debug              green
	newBrush("1;32"), // Trace              green
}

type consoleLogger struct {
	sync.Mutex
	Level    string `json:"level"`
	Colorful bool   `json:"color"`
	LogLevel int
}

func (c *consoleLogger) Init(jsonConfig string) error {
	if len(jsonConfig) == 0 {
		return nil
	}

	err := json.Unmarshal([]byte(jsonConfig), c)
	if runtime.GOOS == "windows" {
		c.Colorful = false
	}

	if l, ok := LevelMap[c.Level]; ok {
		c.LogLevel = l
		return nil
	}

	return err
}

func (c *consoleLogger) LogWrite(when time.Time, msgText interface{}, level int) error {
	if level > c.LogLevel {
		return nil
	}
	msg, ok := msgText.(string)
	if !ok {
		return nil
	}
	if c.Colorful {
		msg = colors[level](msg)
	}
	c.printlnConsole(when, msg)
	return nil
}

func (c *consoleLogger) Destroy() {

}

func (c *consoleLogger) printlnConsole(when time.Time, msg string) {
	c.Lock()
	defer c.Unlock()
	os.Stdout.Write(append([]byte(msg), '\n'))
}

func init() {
	Register(AdapterConsole, &consoleLogger{
		LogLevel: LevelDebug,
		Colorful: runtime.GOOS != "windows",
	})
}