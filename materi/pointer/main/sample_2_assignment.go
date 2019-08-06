package main
import "fmt"

func PrimitiveAssignment(){
	var sampleTwoNotPointerToPointer string
	var sampleTwoPointer *string = &sampleTwoNotPointerToPointer
	fmt.Println(sampleTwoPointer)
	fmt.Printf("%p\n", sampleTwoPointer)
	fmt.Println(*sampleTwoPointer)

	*sampleTwoPointer="This Is Pointer Value"
	fmt.Println(*sampleTwoPointer)
}

func StructAssignment(){
	
	type PersonOne struct {
		Name	string
		NotAge	int
	}
	sampleOnePointerStruct := &PersonOne{
		Name: "Lukman",
		NotAge: 30,
	}
	fmt.Printf("%p\n", sampleOnePointerStruct)
	fmt.Println(sampleOnePointerStruct)

	type PersonTwo struct {
		Name	string
		NotAge	int
	}
	sampleTwoPointerStruct := &PersonTwo{}
	sampleTwoPointerStruct.Name = "Lukman"
	sampleTwoPointerStruct.NotAge = 30
	fmt.Printf("%p\n", sampleTwoPointerStruct)
	fmt.Println(sampleTwoPointerStruct)
}