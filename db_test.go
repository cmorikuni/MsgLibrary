package main

import (
    "fmt"
    "strconv"
	"testing"
)

func TestAll(t *testing.T) {
    fmt.Println("\nTestAll")
	open()
	
    mc := queryAll()
    if len(mc) != 4 {
        t.Fail()
    }
    print(mc)
}

func TestSearch(t *testing.T) {
    fmt.Println("\nTestSearch")
    open() 

    mc := querySearch("test")
    if len(mc) != 2 {
        t.Fail()
    }
    print(mc)
}

func TestType(t *testing.T) {
    fmt.Println("\nTestType")
    open()

    mc := queryType("test")
    if len(mc) != 1 {
        t.Fail()
    }
    print(mc)
}

func TestMessageErr(t *testing.T) {
    fmt.Println("\nTestMessage")
    open()

    vars := []string{"a", "b", "c"}
    mc := queryMsg("1", vars)
    fmt.Println(mc)
}

func TestMessage(t *testing.T) {
    fmt.Println("\nTestMessage")
    open()

    vars := []string{"a", "b"}
    mc := queryMsg("1", vars)
    fmt.Println(mc)
}

func print(mc []*MsgCategory) {
    fmt.Println("Categories:", len(mc))

    for _, item := range mc {
        fmt.Println(item.Category + ": " + strconv.Itoa(len(item.Messages)))
        for _, msg := range item.Messages {
            fmt.Println(msg)
        }
    }
}