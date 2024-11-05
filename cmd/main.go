package main

import (
    "fmt"
    "example.com/ranges"
)

func main() {
    empty := ranges.Empty()
    interval1 := ranges.Opened(10, 20)
    interval2 := empty.Intersection(interval1)
    fmt.Print(empty.Equal(interval2))
}
