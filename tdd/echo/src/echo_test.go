package main

import (
    "os"
    "testing"
)

func TestEchoWorks(t *testing.T) {
    args := ["one" "two" "three"]
    result := echo(args)
    if result != "one two three" {
        t.Error("Not matching")
    }
}