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

func findFoodIds() []string {
    var foodIds []string
    doForEachLineExceptFirst(func (line string) {
        elements := strings.Split(line,"$")
        foodIds = append(foodIds, elements[0])
    }, "FOOD_NM.txt")

    return foodIds
}

func findNutrientIds() []string {
    var nutIds []string
    doForEachLineExceptFirst(func (line string) {
        elements := strings.Split(line,"$")
        nutIds = append(nutIds, elements[0])
    }, "NT_NM.txt")

    return nutIds
}

func nutrientName(id string) string {
    var retString string = ""

    doForEachLine(func (line string) {
        elements := strings.Split(line,"$")
        if elements[0] == id {
            retString = elements[4]
        }
    }, "NT_NM.txt")

    if(retString != "") {
        return retString
    }

    return "Not found"
}

func foodName(id string) string {
    var retString string = ""
    doForEachLine(func (line string) {
        elements := strings.Split(line,"$")
        if elements[0] == id {
            retString = elements[4]
        }
    }, "FOOD_NM.txt")
    
    if(retString != "") {
        return retString
    }

    return "Not found"
}

func nutrientAmnts(foodId string) [870]string {
    var retArray [870]string
    doForEachLineExceptFirst(func (line string) {
        elements := strings.Split(line,"$")
        if elements[0] == foodId{
            nid, err := strconv.Atoi(elements[1])
            if err != nil { panic(err) }
            retArray[nid] = elements[2]
        }
    }, "NT_AMT.txt")

    return retArray
}

func main(){
    //find a list of all food ids
    foodIds := findFoodIds()
    //find a list of all nutrient ids
    nutIds := findNutrientIds()
    //generate an array of foods and nutrients in a multidimensional array
    
    //Print the header for the file
    fmt.Print("\"FOOD NAME\"")
    for i := 0 ; i < len(nutIds) ; i++ {
        fmt.Print("$")
        fmt.Print(nutrientName(nutIds[i]))
    }
    fmt.Println("")
    
    for i := 0 ; i < len(foodIds) ; i++ {
        fmt.Print(foodName(foodIds[i]))
        nutrients := nutrientAmnts(foodIds[i])

        for k := 0 ; k < len(nutIds) ; k++ {
            fmt.Print("$")
            nid, err := strconv.Atoi(nutIds[k])
            if  err != nil { panic(err) }
            if nutrients[nid] == "" {
                fmt.Print("0")
            } else {
                fmt.Print(nutrients[nid])
            }
        }
        fmt.Println("")
    }
}
