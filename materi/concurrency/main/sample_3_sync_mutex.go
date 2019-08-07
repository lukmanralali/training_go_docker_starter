package main
import (
	"fmt"
	"sync"
	// "encoding/json"
	"time"
)
// import "math/rand"

type WaitGroupSyncMutex struct{
	
}

func (wgm *WaitGroupSyncMutex) SampleSync() {
	wgm.blackpinkDanceExercise()
}

func (wgm *WaitGroupSyncMutex) blackpinkDanceExercise(){

	// prepare data
	blackpink := []string{"Jisoo", "Rose", "Lisa", "Jennie"}
	dances := []Dance{}
	wgm.addDanceOption(&dances, "model-1",2)
	wgm.addDanceOption(&dances, "model-2",6)
	wgm.addDanceOption(&dances, "model-3",8)
	wgm.addDanceOption(&dances, "model-4",9)
	result := map[string]int{}

	// set environtment
	minimumDanceMovement := 5
	waitGroup := &sync.WaitGroup{}
	mtx := &sync.Mutex{}
	waitGroup.Add(minimumDanceMovement*len(blackpink))
	// do exercise
	for i := 0; i < minimumDanceMovement; i++ {
		// waitGroup.Add(len(blackpink))
		for _, member := range blackpink {
			go wgm.doExercise(waitGroup,mtx,member,dances,result)
		}
	}
	fmt.Println()
	waitGroup.Wait()
	fmt.Println(result)
}

func (wgm *WaitGroupSyncMutex) doExercise(wg *sync.WaitGroup, mtx *sync.Mutex, member string, dances []Dance, result map[string]int){
	defer wg.Done()
	defer mtx.Unlock()
	mtx.Lock()
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

func (wgm *WaitGroupSyncMutex) addDanceOption(dances *[]Dance,name string,level int){
	*dances = append(*dances, Dance{
		Name: name,
		DifficultyLevel: level,
	})
}