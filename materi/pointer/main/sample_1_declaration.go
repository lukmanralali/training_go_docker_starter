package main
import "fmt"

func PrimitiveDeclaration(){
	// value type ---------------------------------------------
	var sampleNotPointerOne string = "Value Sample One"
	fmt.Println(sampleNotPointerOne)

	sampleNotPointerTwo := "Value Sample Two"
	fmt.Println(sampleNotPointerTwo)

	fmt.Println("--")
	// reference type ----------------------------------------
	// declaration
	sampleOnePointer := new(string)
	fmt.Println(sampleOnePointer)
	fmt.Println(*sampleOnePointer)

	var sampleTwoNotPointerToPointer string
	var sampleTwoPointer *string = &sampleTwoNotPointerToPointer
	fmt.Println(sampleTwoPointer)
	fmt.Printf("%p\n", sampleTwoPointer)
	fmt.Println(*sampleTwoPointer)
}

func StructDeclaration(){
	type PersonOne struct {
		Name	string
		NotAge	int
	}
	sampleOnePointerStruct := new(PersonOne)
	fmt.Printf("%p\n", sampleOnePointerStruct)

	type PersonTwo struct {
		Name	string
		NotAge	int
	}
	sampleTwoPointerStruct := &PersonTwo{}
	fmt.Printf("%p\n", sampleTwoPointerStruct)
}