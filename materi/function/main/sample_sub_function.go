package main
import "fmt"

func SampleTwo(){
	variadicSample()
	defferSample()
}

func variadicSample(){

	variadicOne(1, 2, 3)
	numbers := []int{1, 2, 3, 4}
	variadicOne(numbers...)
	variadicTwo("-",numbers...)
}

func variadicOne(numbers ...int) {
    fmt.Print(numbers, " ")
    total := 0
    for _, number := range numbers {
        total += number
    }
    fmt.Println("result:",total)
}

func variadicTwo(action string,numbers ...int) {
    fmt.Println(numbers)
    total := 0
    for _, number := range numbers {
    	switch action {
		    case "+":
		        total += number
		    case "-":
		        total -= number
	    }
    }
    fmt.Println("result:",total)
}

func defferSample(){
	defer fmt.Println("sorry this is just stupid story")
	fmt.Println("dr.strange : lets go fellas, we need to kill thanos")
	fmt.Println("thor : i want his head")
	defer fmt.Println("thanos dead. end. thanks!!!")
	fmt.Println("ironman : you cant chop his head if you dead thor")
	fmt.Println("spiderman : mr.stark, i am alive i dont wanna die again")
	fmt.Println("me : i just need to sit back and see you all die, then i will kil him")
}