# KDTree

An idiomatic go kdtree.

Abstracts away comparisons to user code,
so it's similar to the sort and heap packages.

But, it does better by allowing the user to propogate their own
errors.

The user of the lib must implement:

```go
type Interface interface {
  DistSqr(i, j int) float64
  CompareDimension(i, j int, dim int) float64
  For(func(int) bool)
  Insert(v interface{}) (index int, err error)
  Delete(i int) error
}
```

Which allows for flexible type comparisons, ie. string comparisons can be implemented on the
client side, or int comparisons as long as they are cast to float64

