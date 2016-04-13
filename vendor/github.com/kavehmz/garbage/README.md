# garbage
tests and garbage

```go
func callPointertoFunc() {
	x := new(func(int) bool)
	*x = func(x int) bool { return false }

	fmt.Println((*x)(4))
}

```
