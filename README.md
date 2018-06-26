# Domain Parser for Go

The package domain-parser provides domain name parser based on url.Parse that is golang function

```go 
package main

import (
	"fmt"

	domain "github.com/onerciller/domain-parser"
)

func main() {

	example, _ := domain.Parse("http://example.com") 

	fmt.Println(example) //example.com

	subdomain, _ := domain.Parse("http://subdomain.example.com")

	fmt.Println(subdomain) //example.com

	w, _ := domain.Parse("http://www.example.com")

	fmt.Println(w) //example.com

	wSubdomain, _ := domain.Parse("http://www.subdomain.example.com")

	fmt.Println(wSubdomain) //example.com

}

```
