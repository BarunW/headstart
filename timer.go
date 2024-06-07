package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"golang.org/x/term"
)

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
    duration time.Duration 
    terminalHeight int
    terminalwidth int
    rwMut sync.RWMutex
}

// This function have a side effect when
// user provide invalid unit of time
func NewTimer(duration time.Duration){ 
    t := &Timer{
        duration: duration,
        rwMut: sync.RWMutex{},
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
func(t *Timer) Renderer(d time.Duration){
    ticker := time.NewTicker(16666700 * time.Nanosecond)
    timeOut := time.NewTimer(d)
    var(
       ms uint8
       s uint8
       m uint8
       h uint8
   )

    var f = func(){ 
      if m >= 60{
          h += 1
          m  = 0
      }
      if s >= 60{
          m += 1
          s  = 0
      }
      if ms > 60{
        ms = 0
        s += 1 
      }

    }
    outer:
    for{
        select{
        case <-ticker.C: 
            ms++
            fmt.Printf("  %d%s : %d%s : %d%s : %d\r", h, Green, m, Blue, s, Cyan, ms) 
            f()
        case <-timeOut.C:
            break outer 
        }
    }
    return
}

func(t *Timer) Start(){
    timer := time.NewTimer(t.duration)   
    outer:
    for{
        select {
        case <-timer.C:
            fmt.Println("\nTimeOut")
            break outer
        default: 
            t.Renderer(t.duration)
        }
    }
    os.Exit(1)
}
