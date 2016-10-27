package main

import (
    "fmt"
	"testing"
)

func TestQueryAll(t *testing.T) {
	open()
	mc := queryAll()
	close()
    if len(mc) != 4 {
        t.Fail()
    }
    print(mc)
}

func print(mc []*MsgCategory) {
    fmt.Println("Categories:", len(mc))

    for _, item := range mc {
        fmt.Println(item.Category)
        for _, msg := range item.Messages {
            fmt.Println(msg)
        }
    }
}