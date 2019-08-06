package main
import "fmt"

func SampleOne(){
	zeroParameterZeroReturn()
	manyParameterZeroReturn("i am ironman!", 3000)
	zeroParameterManyReturn()
	manyParameterManyReturn("i am ironman!", 3000)
}

func zeroParameterZeroReturn(){
	fmt.Println("i am nobody")
}

func manyParameterZeroReturn(superHeroWord string, loveMeter int){
	fmt.Println(superHeroWord)
	fmt.Println(loveMeter)
}

func zeroParameterManyReturn() (string, int){
	return "i am ironman", 3000
}

func manyParameterManyReturn(superHeroWord string, loveMeter int) (string, int){
	superHeroWord = "ironman is died"
	loveMeter = 0
	return superHeroWord, loveMeter
}