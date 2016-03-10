package logging

import (
    "log"
    "os"
)

var instance *log.Logger

func getLoggerInstance()(*log.Logger) {
    if instance == nil {
        log.New(os.Stdout, "EVENT:", log.Ldate|log.Ltime|log.Lshortfile)
    }

    return instance
}

func LogEvent(s string) {
    instance.Println(s)
}
