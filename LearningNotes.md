
指针问题
```go
package main

import "fmt"

func main() {
	a := make([]*int, 1)
	testPointer(a[0])
	fmt.Println(a)
}
func testPointer(a *int) {
	*a = 10
}

```
