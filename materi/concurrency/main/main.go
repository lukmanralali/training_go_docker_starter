package main
import "math/rand"

func main() {

    // sampleOne := &BasicGoRoutine{}
    // sampleSyncOne := &WaitGroupSync{}
    // sampleSyncTwo := &WaitGroupSyncMutex{}
    sampleSyncThree := &WaitGroupSyncChannel{}
    // sampleOne.RunExample()
    // sampleSyncOne.SampleSync()
    // sampleSyncTwo.SampleSync()
    sampleSyncThree.SampleSync()
}

func GetRandomInteger(min int,max int) int {
    return (rand.Intn(max - min + 1) + min)
}