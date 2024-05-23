# goword   
[![test](https://github.com/frederikhs/goword/actions/workflows/test.yml/badge.svg)](https://github.com/frederikhs/goword/actions/workflows/test.yml)

Go Package to extract text from word docx files

### usage

```
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
