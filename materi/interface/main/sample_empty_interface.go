package main

import "fmt"

func SampleEmptyInterface(){
    sampleOne()
    sampleTwo()
    sampleThree()
    sampleFour()
}

func sampleOne(){
    var secret interface{}

    secret = "ethan hunt"
    fmt.Println(secret)

    secret = []string{"apple", "manggo", "banana"}
    fmt.Println(secret)

    secret = 12.4
    fmt.Println(secret)
}

func sampleTwo(){
    var sampleSlice []interface{}

    sampleSlice = append(sampleSlice, "Ghost Protocol", "Rogue Nation", "Fallout")
    fmt.Println(sampleSlice)
}

func sampleThree(){
    type MissionImposible struct {
        SeriesName string
        ReleaseYears  int
    }
    var sampleSlice []interface{}

    sampleSlice = append(sampleSlice,
    MissionImposible{
        SeriesName : "Ghost Protocol",
        ReleaseYears: 2011,
    },MissionImposible{
        SeriesName : "Rogue Nation",
        ReleaseYears: 2015,
    })
    fmt.Println(sampleSlice)
}