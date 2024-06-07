package main

import (
	"flag"
	"log/slog"
	"os"
	"strconv"
	"time"
)

func DefineTimerFlag() func()time.Duration{
    s := flag.Bool("s", false, "second flag for timer")
    m := flag.Bool("m", false, "minute flag for timer")
    h := flag.Bool("h", false, "hour flag for timer")

    return func() time.Duration{
            if *s{return time.Second }
            if *m{return time.Minute }
            if *h{return time.Hour }
            return -1
    }
}
func main() {
	args := os.Args
	hs := NewHSCommands()
    
    tf := DefineTimerFlag()
    flag.Parse()
    unit := tf()

    
    if unit != -1{
        dur, err := strconv.Atoi(flag.Arg(0)) 
        if err != nil{
            slog.Error("Inavlid duration")
            os.Exit(1)
        }
        NewTimer(time.Duration(dur) * unit )
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
