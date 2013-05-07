package main

import (
    "fmt"
    "os"
    "io"
    "bufio"
    "strings"
    "flag"
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

func printAmount(amount, foodName string) {
    var foodId string
    var measId string
    var convAmt string
    var measName string

    //determine the food id
    doForEachLineExceptFirst(func (line string) {
        elements := strings.Split(line,"$")
        if(elements[4] == foodName) {
            foodId = elements[0]
        }
    }, "../Sources/Canadian Database/FOOD_NM.txt")

    //find the measure id and conv amt
    first := true
    doForEachLineExceptFirst(func (line string) {
        elements := strings.Split(line,"$")
        if(elements[0] == foodId && first) {
            measId = elements[1]
            convAmt = elements[2]
            first = false
        }
    }, "../Sources/Canadian Database/CONV_FAC.txt")

    //find the measName
    doForEachLineExceptFirst(func (line string) {
        elements := strings.Split(line,"$")
        if(elements[0] == measId) {
            measName = elements[1]
        }
    }, "../Sources/Canadian Database/MEASURE.txt")
    fmt.Println(convAmt + " of " + measName)
}

func main(){
    flag.Parse()
    if (flag.NArg() != 1) {
        fmt.Println("Usage: go run ./whichFood.go [foodNumber]")
    } else {
        i := 0
        arg, err := strconv.Atoi(flag.Arg(0))
        if err != nil { panic(err) }

        doForEachLine(func (line string) {
            elements := strings.Split(line,"$")
            if i == 0 {
                printAmount("1", elements[arg])
                i = 1
            }
        }, "nutrientsMatrixListLowerTranspose")
    }
}
