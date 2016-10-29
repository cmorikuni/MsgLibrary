package main

import (
    "fmt"
    "strconv"
	"testing"
)

// Tests with inserted ID
var insertId string

func TestInsert(t *testing.T) {
    fmt.Println("\nTestInsert")
    open()

    mc := insert("Type", "New Msg %%")
    for _, item := range mc {
        for _, msg := range item.Messages {
            insertId = strconv.Itoa(msg.Oid)
        }
    }
    print(mc)
}

func TestUpdate(t *testing.T) {
    fmt.Println("\nTestUpdate")
    open()    

    mc := update("0", "message", "TESTING")
    print(mc)
}

func TestMessageErr(t *testing.T) {
    fmt.Println("\nTestMessage")
    open()

    vars := []string{"a", "b", "c"}
    mc := queryMsg(insertId, vars)
    fmt.Println(mc)
}

func TestMessage(t *testing.T) {
    fmt.Println("\nTestMessage")
    open()

    vars := []string{"a", "b"}
    mc := queryMsg(insertId, vars)
    fmt.Println(mc)
}


func TestDelete(t *testing.T) {
    fmt.Println("\nTestDelete")
    str := delete(insertId)
    fmt.Println(str)
}
// End

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

func print(mc []*MsgCategory) {
    fmt.Println("Categories:", len(mc))

    for _, item := range mc {
        fmt.Println(item.Category + ": " + strconv.Itoa(len(item.Messages)))
        for _, msg := range item.Messages {
            fmt.Println(msg)
        }
    }
}