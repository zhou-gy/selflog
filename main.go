package main

import(
	"logput/console"
) 

var Logger console.Logger

func TestFunc(){
	Logger.DEBUG("xxxxxxxxxxx")
	Logger.WARN("xxxxxxxxxxx")
	Logger.ERROR("xxxxxxxxxxx %s %d","lalala",100)
}
func main() {
	Logger = Logger.SetLevel("INFO","json")
	// Logger.Ptype = "json"
	
	TestFunc()

	
}
