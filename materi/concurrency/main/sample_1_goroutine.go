package main
import "fmt"
import "time"

type BasicGoRoutine struct{
}

func (groutine *BasicGoRoutine) RunExample(){
    go groutine.iterateAndPrint("process-A")
    go groutine.iterateAndPrint("process-B")
    go groutine.iterateAndPrint("process-C")
    fmt.Scanln()
}

func (groutine *BasicGoRoutine) iterateAndPrint(identifier string) {
    for i := 1; i <= 3; i++ {
        someRandomProccesTime := GetRandomInteger(100,1000)
        time.Sleep(time.Duration(someRandomProccesTime) * time.Millisecond)
        fmt.Println(identifier, ":", i, "| time:",someRandomProccesTime)
    }
}