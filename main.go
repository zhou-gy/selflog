package main

import (
	"logput/console"
)

var logger console.Logger

func TestFunc(){
	logger.INFO("this is info log....")
	logger.DEBUG("xxxxxxxxxxxxx%d", 10)
	logger.WARN("WARNNNNNNNNNNNNNNN")
	logger.ERROR("ERRORRRRRRRRRR")
	logger.FATAL("FATALLLLLLLLLLLLLLL")
}

func main() {
	logger = logger.SetLevel("INFO")
	TestFunc()
}
