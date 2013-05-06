package main

import (
    "fmt"
    "os"
    "io"
    "bufio"
    "strings"
    "strconv"
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

    reader := bufio.NewReaderSize(file, 4*1024)

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

func doForEachLineExceptFirst(do stringOperation, filename string) {
    var firstLine bool = true
    doForEachLine(func (line string) {
        if !firstLine {
            do(line)
        } else {
            firstLine = false
        }
    }, filename)
}

func main(){
    //The first thing that the lp_solve file expects is 
    //an objective function:
    fmt.Println("min: ;")

    var numberOfVariables int

    //Next we print out the nutrition constraints:
    doForEachLineExceptFirst(func (line string) {
        elements := strings.Split(line,"$")
        for i := 0 ; i < len(elements) - 1 ; i++ {
            if i > 0 {
                fmt.Print(" + " + elements[i] + "x" + strconv.Itoa(i + 1))
            } else {
                fmt.Print(elements[i] + "x" + strconv.Itoa(i + 1))
            }
        }
        fmt.Println(" <= " + elements[len(elements) - 1] + ";")
        numberOfVariables = len(elements)
    }, "NutrientMatrixUpper")
    
    doForEachLineExceptFirst(func (line string) {
        elements := strings.Split(line,"$")
        for i := 0 ; i < len(elements) - 1 ; i++ {
            if i > 0 {
                fmt.Print(" + " + elements[i] + "x" + strconv.Itoa(i + 1))
            } else {
                fmt.Print(elements[i] + "x" + strconv.Itoa(i + 1))
            }
        }
        fmt.Println(" >= " + elements[len(elements) - 1] + ";")
        numberOfVariables = len(elements)
    }, "NutrientMatrixLower")

    //Now we print out the greater than zero constraints:
    for i := 0 ; i < numberOfVariables ; i++ {
        fmt.Println("x" + strconv.Itoa(i + 1) + " >= 0;")

    }
}
