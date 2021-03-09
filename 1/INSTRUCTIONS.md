# Installation and First Program

## Download Installer

<https://golang.org/dl/>

## Create a module

```bash
$ mkdir hello
$ cd hello
$ go mod init github.com/rjha1/hello
```

### Generated go.mod

```bash
$ cat hello/go.mod
module github.com/rjha1/hello

go 1.16
```

## Hello World

```go
// hello/hello.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## Build and Run

```bash
$ go build hello.go
$ ./hello
```

```bash
$ go run hello.go
```
