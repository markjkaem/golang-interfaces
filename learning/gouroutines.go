package learning

import (
	"fmt"
	"sync"
	"time"
)

var mx = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func Goroutine() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are: %v", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("\nThe result from the database is: ", dbData[i])
	save(dbData[i])
	logging()
	wg.Done()
}

func save(result string) {
	mx.Lock()
	results = append(results, result)
	mx.Unlock()
}

func logging() {
	mx.RLock()
	fmt.Printf("\nthe current results are%v", results)
	mx.RUnlock()
}
