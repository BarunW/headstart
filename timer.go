package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"golang.org/x/term"
)

type TimeUnit string


const (
	Reset  = "\033[0m"
	Blue   = "\033[34;1m"
	Red    = "\033[31m"
	Green  = "\033[3221m"
	White  = "\033[37;1m"
	Yellow = "\033[33;1m"
	Purple = "\033[35m"
    Cyan   = "\033[36;1m"
	Gray   = "\033[90m"

    // 256 color pallete
    Yellow256 = "\033[;38;5;214m"
    Blue256   = "\033[;38;5;26m"
	Cyan256   = "\033[;38;5;51m"
    Orange256 = "\033[;38;5;208m"
)

type Timer struct{
	timeUnitAndDuration map[TimeUnit]uint
    terminalHeight int
    terminalwidth int
    rwMut sync.RWMutex
}

// This function have a side effect when
// user provide invalid unit of time
func NewTimer(timeUintAndDuration map[TimeUnit]uint){ 
    t := &Timer{
        rwMut: sync.RWMutex{},
		timeUnitAndDuration: timeUintAndDuration,
    }
    t.Start()
}

func(t *Timer) HandleScreenRes(){    
    for{
        w, h, err := term.GetSize(int(os.Stdin.Fd()))     
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error getting terminal size: %v\n", err)
        }
        t.rwMut.Lock()
        t.terminalHeight = h
        t.terminalwidth  = w
        t.rwMut.Unlock()
    }

}
func(t *Timer) Renderer() context.Context{
	ticker := time.NewTicker(16666700 * time.Nanosecond)
	ctx, cancel := context.WithCancel(context.Background())
	var(
		frame uint 
		s uint = 0
		m uint = 0
		h uint = 0
	)
	go func(){
		outer:
		for{
			select{
				case <-ticker.C: 
				frame++
				if m == 60{
					h += 1
					m  = 0
				}
				if s == 60{
					m += 1
					s  = 0
				}
				if frame == 60{
					frame = 0
					s += 1 
				}
				
				fmt.Printf("%sThis is not the precise to millisecond: [  %d%s : %d%s : %d%s : %d ]\r ", Red, h, Green, m, Blue, s, Cyan, frame) 
				if  h == t.timeUnitAndDuration["h"] && 
				m == t.timeUnitAndDuration["m"] && 
				s == t.timeUnitAndDuration["s"]{
					break outer
				}
			}
		}
		fmt.Printf("\n")
		cancel()		
	}()
    return ctx
}

func(t *Timer) Start(){
	ctx := t.Renderer()
	<-ctx.Done()
}
