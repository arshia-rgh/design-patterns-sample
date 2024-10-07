package creational

import (
	"fmt"
	"sync"
)

type Single struct{}

var lock = &sync.Mutex{}

var singleInstance *Single

func getInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating instance")
			singleInstance = &Single{}
		} else {
			fmt.Println("Instance already exists")
		}
	} else {
		fmt.Println("Instance already exists")
	}

	return singleInstance
}

//func main() {
//	for i := 0; i < 10; i++ {
//		go getInstance()
//	}
//
//	fmt.Scanln()
//}
