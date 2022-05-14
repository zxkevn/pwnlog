package main 

import "pwnlog"

func main() {
    logger := pwnlog.New(pwnlog.DebugLevel)
    logger.Debug("this is debug output!")

    logger.SetLevel(pwnlog.InfoLevel)
    logger.Debug("you should NOT see this debug output")
    logger.Info("changed log level to %s", "info")
    logger.Warning("warning: error code %d", 42)
    logger.Error("ERROR!")
    
    // goodbye program
    logger.Fatal("F A T A L")
}
