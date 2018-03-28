## Using goroutine for slice iterating
---
Currently benchmark is odd, goroutine is slower than directly.

```
$ go test -bench=. | grep ns
    1000           1201456 ns/op
    1000           1084547 ns/op
```