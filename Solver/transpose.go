package main

import (
    "fmt"
    "os"
    "io"
    "flag"
    "bufio"
    "strings"
)

type stringOperation func(string)

func doForEachLine(do stringOperation, filename string) {
    //open the file
    file, err := os.Open(filename)
    if err != nil { panic(err) }

    defer func() {
        if err := file.Close(); err != nil {
            panic(err)
        }
    }()

    reader := bufio.NewReaderSize(file, 4*64*1024)

    var endof bool = false
    for !endof {
        line, isPrefix, err := reader.ReadLine()
        if err != nil && err != io.EOF { panic(err) }
        if err == io.EOF {
            endof = true
        } else {
            //do stuff to line
            do(string(line))
        }

        if isPrefix {
            fmt.Println("buffer not large enough")
            return
        }
    }
}

func main(){
    //Check to make sure that the program is being used correctly:
    flag.Parse()
    if flag.NArg() != 1 {
        fmt.Println("Usage: go run ./transpose.go [filename]")
    } else {

        //First, we need to go through each line of the file, and load the data into a multidimensional string array
        //fileContents[foodId][nutId]
        var fileContents [][]string = nil
        doForEachLine(func (line string) {
            elements := strings.Split(line,"$")
            fileContents = append(fileContents, elements)
        }, flag.Arg(0))

        for k := 0 ; k < len(fileContents[0]) ; k++ {
            for i := 0 ; i < len(fileContents); i++ {
                fmt.Print(fileContents[i][k])
                if (i < len(fileContents) - 1) { fmt.Print("$") }
            }
            fmt.Println("")
        }
    }
}
