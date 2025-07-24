package main

import (
	"flag"
	"os"
)

func DefineTimerFlag() func() map[TimeUnit]uint{
    s := flag.Uint("s", 0, "second flag for timer")
    m := flag.Uint("m", 0, "minute flag for timer")
    h := flag.Uint("h", 0, "hour flag for timer")

    return func() map[TimeUnit]uint{
		return map[TimeUnit]uint{
			"h" : *h, 
			"m" : *m, 
			"s" : *s,
		}
	}
}
func main() {
	args := os.Args
	hs := NewHSCommands()
    
    tf := DefineTimerFlag()
    flag.Parse()
    timeUintsAndDuration := tf()
	
	shouldRunTimer := false
	for _, val := range timeUintsAndDuration{
		if val != 0 { shouldRunTimer = true } 
	}
	
	if shouldRunTimer { 
		NewTimer(timeUintsAndDuration)	
		return
	}
				

	if len(os.Args) >= 3 {
		hs.processCommandWithSubcmd(Command(args[1]), args[2:]...)
		return
	}

    if len(os.Args) == 2{
        hs.LinkOutPut(os.Args[1])        
        return
    }

	hs.ShowAllCommands()
}
