package main

import (
	compiler "aolda_node/compiler"
	database "aolda_node/database"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		var a []string
		a = append(a, "1")
		a = append(a, "2")
		compiler.ExecuteJS("add", "add", a)

		// contract.ListenEvent()
	}()

	go func() {
		database.Boltdb()
		defer wg.Done()
	}()

	wg.Wait()
}
