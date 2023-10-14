package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	userName := fetchUser()
	response := make(chan any, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go fetchUserLikes(userName, response, wg)
	go fetchUserMatch(userName, response, wg)

	wg.Wait()
	close(response)

	for resp := range response {
		fmt.Println("resp: ", resp)
	}

	fmt.Println("took: ", time.Since(start))

}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "hillheim"
}

func fetchUserLikes(userName string, responseChannel chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)
	responseChannel <- 11
	wg.Done()
}

func fetchUserMatch(userName string, responseChannel chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	responseChannel <- "jane"
	wg.Done()
}
