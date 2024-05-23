# goword   
[![test](https://github.com/frederikhs/goword/actions/workflows/test.yml/badge.svg)](https://github.com/frederikhs/goword/actions/workflows/test.yml)

Go Package to extract text from word docx files

### usage

```
import "github.com/frederikhs/goword"

func main() {
    text, err := goword.ParseText("1.docx")
    if err != nil {
        log.Panic(err)
    }
    fmt.Println(text)
}
```
