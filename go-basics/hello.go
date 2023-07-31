 
package main

import "fmt"



func main () {

    fmt.Println("hello go")

    var a int = 5
    var b float32 = 4.32
    const pi float64 = 3.1415139475

    var (
        c = 8
        d = 7
    )

    e,f := 9,10

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(pi)
    fmt.Println(c)
    fmt.Println(d)
    fmt.Println(e+f)

    isbool := true
    isnot := false

    fmt.Println(isbool && isnot)
    fmt.Println(&isbool)




}
