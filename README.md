### Golexer

A simple golnag library/utility to tokenize a file into individual tokens usually for parsing scripts/programming languages.

Please note that this utility for now cannot be used to parse simple text files containing only plain text and is more appriopriate 
for lexing syntax based languages/scripts.


### Usage

```
make build

./target/golexer path/to/file

```


### Example

A file like :

```
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dropdevrahul/golexer/golexer"
)

func main() {
	usage := `Usage: golexer FILE`
	//save := flag.Bool("c", false, "Save results to a file instead of printing")

	if len(os.Args) < 2 {
		fmt.Println(usage)
		return
	}

	tz := golexer.NewTokenizer()

	tokens, err := tz.LexFile(os.Args[1])
	if err != nil {
		log.Panic(err)
	}
	for _, t := range tokens {
		fmt.Printf("%s\n", t.Value)
	}
	fmt.Println()
}
```

will be parse into

```
package
main
import
(
"fmt"
"log"
"os"
"github.com/dropdevrahul/golexer/golexer"
)
func
main
(
)
{
usage
:=
`Usage:
golexer
FILE`
//save
:=
flag.Bool
(
"c"
,
false
,
"Save results to a file instead of printing"
)
if
len
(
os.Args
)
<
2
{
fmt.Println
(
usage
)
return
}
tz
:=
golexer.NewTokenizer
(
)
tokens
,
err
:=
tz.LexFile
(
os.Args
[
1
]
)
if
err
!=
nil
{
log.Panic
(
err
)
}
for
_
,
t
:=
range
tokens
{
fmt.Printf
(
"%s\n"
,
t.Value
)
}
fmt.Println
(
)
}

```

