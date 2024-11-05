## Ranges

Range is just a datastructure that let you manipulate intervals or ranges and performs operation between ranges

```go
def main() {
    empty := ranges.Empty()
    interval1 := ranges.Opened(10, 20)
    interval2 := empty.Intersection(interval)
    fmt.Print(empty.Equal(interval2))
}
```

## Thanks

+ Python [portion](https://pypi.org/project/portion/) package
