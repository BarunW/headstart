package main

import (
	"fmt"
	"strings"
	"github.com/eiannone/keyboard"
)

type Render struct{}

func NewRender() *Render{
    return &Render{}
}

func(r *Render) clear(l int){ 
    for i := 0; i <= l; i++{
        fmt.Printf("\033[A")
        fmt.Print("\033[K")
    }
}

func( r*Render) appendTo(v map[string]string)[][]string {
    container := [][]string{}
    for key, value := range v{
        container = append(container, []string{"", key, value}) 
    }
    return container
}

 //func(r *Render) Render(values any){
 //    // can be optimized how i display text
 //    // there is blinking becuase of rendering the whole content when moving
 //    window := r.appendTo(values.(map[string]string))
 //    selector := ">"    
 //    var X int = len(window) - 1
 //    const Y int = 0
 //    // put the selector at position window[X][Y]
 //    window[X][Y] = selector
 //    
 //    // key events chan
 //    event := make(chan keyboard.KeyEvent, 6)
 //    // Print the cursor position.
 //    fmt.Println("Use i and j to move and select")
 //    outer:
 //    for {
 //        for _, value := range window{
 //            fmt.Println(strings.Join(value," "))
 //        }
 //        go HandleKeyInput(event)
 //        e := <-event
 //        if X < len(window) - 1 && e.Rune == 'j'{
 //            window[X][Y] = ""
 //            window[X+1][Y] = selector
 //            X++
 //        }else if X > 0 && e.Rune == 'k'{
 //            window[X][Y] = ""
 //            window[X-1][Y] = selector
 //            X--
 //        } else if e.Key == keyboard.KeyEnter{
 //            break outer
 //        } else if e.Key == keyboard.KeyEsc{
 //            fmt.Println("Exiting")
 //            return
 //        }
 //        r.clear(len(window))
 //    }
 //
 //    s := window[X][1]
 //    fmt.Println(strings.EqualFold(s, "GO"))
 //
 //}
 //

func(r *Render) Render(values any){
    // can be optimized how i display text
    // there is blinking becuase of rendering the whole content when moving
    window := r.appendTo(values.(map[string]string))
    selector := ">"    
    var X int = len(window) - 1
    const Y int = 0
    // put the selector at position window[X][Y]
    window[X][Y] = selector
    
    // key events chan
    event := make(chan keyboard.KeyEvent, 3)
    // Print the cursor position.
    fmt.Println("Use i and j to move and select")
    for _, value := range window{
        fmt.Println(strings.Join(value," "))
    }
    fmt.Print("\033[F")
    
    for {
        go HandleKeyInput(event)
       e := <-event
//          ================================================
//          "/r" is to bring cursor at the  begining of line
//          ================================================
        if X < len(window) - 1 && e.Rune == 'j'{

            window[X][Y] = ""
            fmt.Printf("\r\033[K%s", strings.Join(window[X]," "))
            fmt.Printf("\033[B\r")
            X++
            // X becomes X+1
            window[X][Y] = selector
            fmt.Printf("\r\033[K%s\r", strings.Join(window[X]," "))
            
        }else if X > 0 && e.Rune == 'k'{
    
            window[X][Y] =""
            fmt.Printf("\033[K%s",strings.Join(window[X], " "))
            fmt.Printf("\033[F")
            X--
            window[X][Y] = selector
            fmt.Printf("\033[K%s\r", strings.Join(window[X]," "))

        } else if e.Key == keyboard.KeyEnter{
            fmt.Printf("\033[B\r")
            break

        } else if e.Key == keyboard.KeyEsc{
            fmt.Println("Exiting")
            return
        }
    }
    r.clear(X+1)
    s := window[X][1]
    fmt.Println(s)
}

func HandleKeyInput(c chan <- keyboard.KeyEvent){
    keyEvents, err := keyboard.GetKeys(10)
    if err != nil{
        panic("err")
    }

    defer func() {
        _ = keyboard.Close()
    }()

    event := <-keyEvents
    if event.Err != nil{
        panic(event.Err)
    }
    
    c <-event
}

