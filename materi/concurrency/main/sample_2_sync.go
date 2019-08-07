package main
import (
	"fmt"
	"sync"
	"time"
)
type WaitGroupSync struct{
	
}

type Dance struct{
	Name string
	DifficultyLevel int
}

func (wgs *WaitGroupSync) SampleSync() {
	wgs.blackpinkDanceExercise()
}

func (wgs *WaitGroupSync) blackpinkDanceExercise(){

	// prepare data
	blackpink := []string{"Jisoo", "Rose", "Lisa", "Jennie"}
	dances := []Dance{}
	wgs.addDanceOption(&dances, "model-1",2)
	wgs.addDanceOption(&dances, "model-2",6)
	wgs.addDanceOption(&dances, "model-3",8)
	wgs.addDanceOption(&dances, "model-4",9)
	result := map[string]int{}

	// set environtment
	minimumDanceMovement := 5
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(minimumDanceMovement*len(blackpink))
	// do exercise
	for i := 0; i < minimumDanceMovement; i++ {
		// waitGroup.Add(len(blackpink))
		for _, member := range blackpink {
			go wgs.doExercise(waitGroup,member,dances,result)
		}
	}
	fmt.Println()
	waitGroup.Wait()
	fmt.Println(result)
}

func (wgs *WaitGroupSync) doExercise(wg *sync.WaitGroup, member string, dances []Dance, result map[string]int){
	defer wg.Done()
	index := GetRandomInteger(0,2)
	time.Sleep(time.Duration(dances[index].DifficultyLevel) * time.Second)
	_, ok := result[member]
	if !ok {
		result[member]=0
	}
	result[member] += dances[index].DifficultyLevel
	fmt.Println("member: ",member, " | ",dances[index].DifficultyLevel)
	fmt.Println("result: ",result)
	fmt.Println()
}

func (wg *WaitGroupSync) addDanceOption(dances *[]Dance,name string,level int){
	*dances = append(*dances, Dance{
		Name: name,
		DifficultyLevel: level,
	})
}