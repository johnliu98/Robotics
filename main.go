package main

import (
    "fmt"
    "github.com/johnliu98/robotics/sys"
)

func main()  {
    s := sys.NewSystem(2, 1)
    fmt.Println(s)
}
