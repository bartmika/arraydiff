# arraydiff
[![GoDoc](https://godoc.org/github.com/gomarkdown/markdown?status.svg)](https://pkg.go.dev/github.com/bartmika/arraydiff)
[![Go Report Card](https://goreportcard.com/badge/github.com/bartmika/arraydiff)](https://goreportcard.com/report/github.com/bartmika/arraydiff)
[![License](https://img.shields.io/github/license/bartmika/arraydiff)](https://github.com/bartmika/arraydiff/blob/master/LICENSE)
![Go version](https://img.shields.io/github/go-mod/go-version/bartmika/arraydiff)

Convenience functions to help you find content differences in your arrays. Fully tested and no external dependencies.

## Installation

In your Golang project, please run:

```
go get github.com/bartmika/arraydiff
```

## Documentation

All [documentation](https://pkg.go.dev/github.com/bartmika/arraydiff) can be found here.

## Example Usage

```go
package main

import (
    "fmt"

    "github.com/bartmika/arraydiff"

    func main(){
        // Sample data.
    	oldArr := []uint64{1, 2, 3, 4, 5}
    	newArr := []uint64{1, 2, 6, 7, 8, 9, 10}

    	// See what differences between the two arrays.
    	addedArr, sameArr, removedArr := Uints64(oldArr, newArr)

    	// Correct results shown in comments.
        fmt.Println(addedArr)   // 6, 7, 8, 9, 10
        fmt.Println(sameArr)    // 1, 2
        fmt.Println(removedArr) // 3, 4, 5
    }
)
```

## Contributing

Found a bug? Want a feature to improve your developer experience when finding the difference? Please create an [issue](https://github.com/bartmika/arraydiff/issues).

## License
Made with ❤️ by [Bartlomiej Mika](https://bartlomiejmika.com).   
The project is licensed under the [ISC License](LICENSE).
