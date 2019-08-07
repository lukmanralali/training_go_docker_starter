package main
import "fmt"
import "math/rand"

func PrimitiveEffect(){
	notPointerSampleEffect()
	pointerSampleEffect()
	anotherPointerSampleEffect()
}

func notPointerSampleEffect(){
	var sampleTwoNotPointerToPointer string = "This Is not Pointer Value"
	fmt.Println(sampleTwoNotPointerToPointer)
	changeNotPointerValueOne(sampleTwoNotPointerToPointer)
	fmt.Println(sampleTwoNotPointerToPointer)
	sampleTwoNotPointerToPointer = changeNotPointerValueTwo(sampleTwoNotPointerToPointer)
	fmt.Println(sampleTwoNotPointerToPointer)
}

func changeNotPointerValueOne(payload string){
	payload = "Lets Try To Change - 1"
}

func changeNotPointerValueTwo(payload string) string {
	return "Lets Try To Change - 2"
}

func pointerSampleEffect(){
	var sampleTwoNotPointerToPointer string
	sampleTwoPointer := &sampleTwoNotPointerToPointer
	*sampleTwoPointer = "This Is Pointer Value"
	sampleThreePointer := sampleTwoPointer
	fmt.Println(*sampleTwoPointer)
	changePointerValueOne(sampleTwoPointer)
	fmt.Println(*sampleTwoPointer)
	changePointerValueOne(sampleThreePointer)
	fmt.Println(*sampleTwoPointer)
}

func changePointerValueOne(payload *string){
	*payload = "Lets Try To Change Pointer - "+randStringBytes(5)
}

func randStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func anotherPointerSampleEffect(){
	var valueOne, valueTwo int
	valueOne, valueTwo = 1, 2
	valueOnePointer := &valueOne
	valueTwoPointer := &valueTwo
	fmt.Printf("%d %d\n", *valueOnePointer, *valueTwoPointer)
	swapNumber(valueOnePointer, valueTwoPointer)
	fmt.Printf("%d %d\n", *valueOnePointer, *valueTwoPointer)
}

func swapNumber(valueOne *int, valueTwo *int){
	*valueOne = *valueOne + *valueTwo
	*valueTwo = *valueOne - *valueTwo
	*valueOne = *valueOne - *valueTwo
}