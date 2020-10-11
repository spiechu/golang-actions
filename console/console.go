package console

import (
    "fmt"
    "io"
)

func PrintHelloWorld(w io.Writer) {
    fmt.Fprintf(w, "Hello GitHub Actions!")
}
