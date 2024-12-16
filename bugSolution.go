```go
package main

import "fmt"

func main() {
    var a [5]int
    for i := 0; i < len(a); i++ {
        a[i] = i * 2
    }
    fmt.Println(a)
}
```