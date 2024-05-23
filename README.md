# goword

[![Release](https://img.shields.io/github/v/release/frederikhs/goword.svg)](https://github.com/frederikhs/goword/releases/latest)
[![Quality](https://goreportcard.com/badge/github.com/frederikhs/goword)](https://goreportcard.com/report/github.com/frederikhs/goword)
[![test](https://github.com/frederikhs/goword/actions/workflows/test.yml/badge.svg)](https://github.com/frederikhs/goword/actions/workflows/test.yml)
[![License](https://img.shields.io/github/license/frederikhs/goword)](LICENSE)
[![GoDoc](https://godoc.org/github.com/frederikhs/goword?status.svg)](https://godoc.org/github.com/frederikhs/goword)

Go Package to extract text from word docx files

### usage

```go
package main

import (
    "fmt"
    "github.com/frederikhs/goword"
    "log"
)

func main() {
    text, err := goword.ParseText("1.docx")
    if err != nil {
        log.Panic(err)
    }
    fmt.Println(text)
}
```
