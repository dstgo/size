# size
simple conversion between size and string

## install
```bash
go get -u github.com/dstgo/size@latest
```

## units
```go
type Unit = int

const (
	B Unit = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
)
```

## usage
```go
package main

import (
	"fmt"
	"github.com/dstgo/size"
)

func main() {
	s1 := size.NewInt(1, size.KB)
	fmt.Println(s1)
	s2 := s1.To(size.MB)
	fmt.Println(s2)

	s3, ok := size.Lookup("1.2MB")
	if !ok {
		panic("failed to lookup")
	}
	fmt.Println(s3)
	lookupTo, ok := size.LookupTo("1.2MB", size.KB)
	if !ok {
		panic("failed to lookup")
	}
	fmt.Println(lookupTo)
}
```
output
```
1KB
0.001MB 
1.2MB   
1228.8KB
```