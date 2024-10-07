package creational

import (
	"fmt"
	"sync"
)

type Single struct{}

var lock = &sync.Mutex{}

var singleInstance *Single

func getInstance() *Single {
	if singleInstance == nil { // check if is not nil avoid doing lock operations
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil { // check again if another go routine created one
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

//--------------------------------------------------------

// Singleton 2

var once sync.Once
var singleInstance2 *Single

func getInstance2() *Single {
	if singleInstance2 == nil {
		once.Do(func() {
			fmt.Println("Creating Instance")
			singleInstance2 = &Single{}
		})
	} else {
		fmt.Println("Instance already exists")
	}

	return singleInstance2
}

//func main() {
//	for i := 0; i < 10; i++ {
//		go getInstance()
//	}
//
//	fmt.Scanln()
//}
