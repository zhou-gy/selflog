package main

import (
	"logput/console"

)

var Logger console.Logger

func TestFunc(){
	Logger.INFO("this is info log....")
	Logger.DEBUG("xxxxxxxxxxxxx%d", 10)
	Logger.WARN("WARNNNNNNNNNNNNNNN")
	Logger.ERROR("ERRORRRRRRRRRR")
	Logger.FATAL("FATALLLLLLLLLLLLLLL")
}

func main() {
	Logger = Logger.SetLevel("INFO")
	TestFunc()

}
