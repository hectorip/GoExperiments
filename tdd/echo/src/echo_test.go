package main

import (
    "os"
    "testing"
)

func TestEchoWorks(t *testing.T) {
    args := make([]string, 3)
    args[0] = "one"
    args[1] = "two"
    args[2] = "three"
    result := echo(args)
    if result != "one two three" {
        t.Error("Not matching")
    }
}