package main

import (
    "fmt"
    "os"
    "io"
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

func main(){
    //find a list of all food ids
    foodIds := findFoodIds()
    //find a list of all nutrient ids
    nutIds := findNutrientIds()
/*
    
    //Print the header for the file
    printf("FoodName")
    for nutrientId in nutrientIds {
        printf("$")
        printf(nutrientName(nutrientId))
    }
    printf("\n")

    for in foodIDs {
        printf(foodName)
        for in nutrientIds {
            printf("$")
            printf(nutrientAmnt(nutrientId,foodId))
        }
        printf("\n")
    }*/
}
