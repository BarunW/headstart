package assert

import "fmt"

func Assert(msg string, scope string, isError bool){
    if isError{
        fmt.Printf("[ERROR] %s [SCOPE]: %s\n", msg, scope )
    }
    fmt.Printf("[ASSERT] %s [SCOPE]: %s\n", msg, scope )
}
