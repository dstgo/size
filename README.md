# size
Fast conversion between file size and string



## Install

```sh
$ go get https://github.com/dstgo/size
```



## Use

**conver string to size**

```go
fmt.Println(ParseSize("1.2MB"))
fmt.Println(ParseSize("2.3b"))
fmt.Println(ParseSize("9.0123TB"))

fmt.Println(ParseSize("abcd"))
fmt.Println(ParseSize("MB"))
fmt.Println(ParseSize("KB"))
fmt.Println(ParseSize("Bk"))

fmt.Println(ParseSize(""))
```

output

```
{1.2 1048576}
{2.3 1}
{9.0123 1099511627776}
{0 0}
{0 0}
{0 0}
{0 0}
{0 0}
```



**conver string t target size**

```go
fmt.Println(ParseTargetSize("1.221MB", KB))
fmt.Println(ParseTargetSize("2.12gB", GB))
```

output

```
{1250.304 1024}
{2.12 1073741824}
```



**convert size to string**

```go
size := ParseTargetSize("1.221MB", KB)
fmt.Println(size.TosString())
size1 := ParseTargetSize("1231GB", MB)
fmt.Println(size1.TosString())
```

output

```
1250.30KB
1260544.00MB
```

