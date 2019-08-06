package main
import "fmt"

type Person struct {
	Name	string
	NotAge	int
}

func StructEffect(){	
	sampleOnePointerStruct := &Person{
		Name: "Lukman",
		NotAge: 30,
	}
	fmt.Printf("%p\n", sampleOnePointerStruct)
	fmt.Println(sampleOnePointerStruct)
	changePlayer(sampleOnePointerStruct)
	fmt.Println(sampleOnePointerStruct)
	changePlayerAgain(sampleOnePointerStruct)
	fmt.Println(sampleOnePointerStruct)
}

func changePlayer(player *Person){
	player.Name = "Dumbass"
	player.NotAge = 17
	fmt.Printf("%p\n", player)
}

func changePlayerAgain(player *Person){
	player = &Person{
		Name : "Douche Bag",
		NotAge : 18,
	}
	fmt.Printf("%p\n", player)
}