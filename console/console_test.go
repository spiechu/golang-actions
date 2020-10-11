package console_test

import (
    "../console"
    "bytes"
    "testing"
)

func TestExists(t *testing.T) {
    t.Run("Test hello world print", func(t *testing.T) {
        buffer := &bytes.Buffer{}

        console.PrintHelloWorld(buffer)

        got := buffer.String()
        want := "Hello GitHub Actions!"

        if got != want {
            t.Errorf("Expected '%s', got '%s'", want, got)
        }
    })
}
