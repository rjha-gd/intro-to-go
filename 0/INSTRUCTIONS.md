# Installation

## Download installer or source code

<https://golang.org/dl/>

## Create a module

```bash
mkdir hello
cd hello
go mod init github.com/rjha1/hello
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

## Build and run

```bash
go build hello.go
./hello
```

```bash
go run hello.go
```
