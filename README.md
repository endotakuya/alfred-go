# alfred-go
Alfred 3 workflows library for Go

## Example

First, install `alfred-go`

```shell
$ go get github.com/endotakuya/alfred-go
```

Usage

```main.go
package main

import (
	"github.com/endotakuya/alfred-go"
)

func main() {
	a := alfred.New()

	item := alfred.NewItem()
	item.Title = "example"
	item.Arg = "http://example.com/"
	a.Append(item)

	a.Print()
}
```

and, `go build`


```shell
$ go build -o workflow main.go
```

In the Script Filter's Script box.
(Language: /bin/bash)

```shell
./workflow
```

