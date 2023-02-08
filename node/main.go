package main

import (
	database "aolda_node/database"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//wg.Add(2)
	wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	contract.ListenEvent()
	// }()

	go func() {
		database.Boltdb()
		defer wg.Done()
	}()

	wg.Wait()
}
