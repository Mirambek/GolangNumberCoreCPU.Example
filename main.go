package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"
)

func main() {
	numberCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numberCores)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	scanner := bufio.NewScanner(os.Stdin)
	var processedList = make(map[int]int)
	var result = make([]string, 0)
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	url := make(chan string, numberCores)
	counter := 0
	for ii := 0; ii < numberCores; ii++ {
		wg.Add(1)
		go func(urlCh <-chan string, id int) {
			defer wg.Done()
			for {
				url, more := <-urlCh
				if more == false {

					return
				}
				start := time.Now()
				resp, err := http.Get(url)
				mutex.Lock()
				processedList[id]++
				mutex.Unlock()
				if err != nil {
					mutex.Lock()
					result = append(result, err.Error())
					mutex.Unlock()
					return
				}
				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)
				secs := time.Since(start).Seconds()
				mutex.Lock()
				result = append(result, fmt.Sprintf("%s;%d;%d;%.2fms", url, resp.StatusCode, len(body), secs*1000))
				mutex.Unlock()
			}
		}(url, ii)
	}
S:
	for scanner.Scan() {
		select {
		case <-interrupt:
			break S
		default:
			urlText := scanner.Text()
			if len(urlText) == 0 {
				break S
			}
			url <- urlText
			counter++
		}
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
	close(url)
	wg.Wait()
	for _, v := range result {
		fmt.Printf("%s\n", v)
	}
	fmt.Println("<порядковый номер паралельного потока запросов>:<число запросов>")
	for k, v := range processedList {
		fmt.Printf("%d:%d\n", k, v)
	}
}
