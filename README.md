## Ranges

Range is just a datastructure that let you manipulate intervals or ranges and performs operation between ranges

```go
func main() {
	empty := ranges.Empty()
	interval1 := ranges.Opened(10, 20)
	interval2 := empty.Intersection(interval1)
	fmt.Println(empty.Equal(interval2))
    for v := range interval1.Iter(1) {
        fmt.Println(v)
    }
}
```

## Thanks

+ Python [portion](https://pypi.org/project/portion/) package
