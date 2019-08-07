package main
import (
	"fmt"
	"time"
)


type WaitGroupSyncChannel struct{
}

func (wgm *WaitGroupSyncChannel) SampleSync() {
	wgm.GetData()
}

func (wgm *WaitGroupSyncChannel) GetData(){
	sampleChannel := make(chan string)

	go wgm.getMessage(sampleChannel, "c", "text-c")
	go wgm.getMessage(sampleChannel, "a", "text-a")
	go wgm.getMessage(sampleChannel, "b", "text-b")
	
	var message = <-sampleChannel
	fmt.Println(message)
	wgm.tellAudienceWeFoundData(message)

	var message2 = <-sampleChannel
	fmt.Println(message2)

	var message3 = <-sampleChannel
	fmt.Println(message3)
}

func (wgm *WaitGroupSyncChannel) tellAudienceWeFoundData(data string){
	fmt.Println("We Found Fastest Procces. The Value Is: ",data)
}

func (wgm *WaitGroupSyncChannel) getMessage(sampleChannel chan string, condition string, payload string){
	switch condition {
	case "a":
		time.Sleep(1 * time.Second)
	case "b":
		time.Sleep(5 * time.Second)
	default:
		time.Sleep(10 * time.Second)
	}
	sampleChannel <- payload
}