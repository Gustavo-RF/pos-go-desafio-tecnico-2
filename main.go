package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Results struct {
	Status int `json:"status"`
	Qty    int `json:"qty"`
}

func someTask(id int, url string, data chan int, wg *sync.WaitGroup, m *sync.Mutex, results *[]Results) {
	for taskId := range data {
		fmt.Printf("%s - Worker: %d executed Task %d\n", time.Now(), id, taskId)

		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			panic(errors.New("error while format address"))
		}

		req.Header.Set("Accepts", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		fmt.Printf("Status: %d\n", resp.StatusCode)

		// precisa tratar condicao de corrida
		m.Lock()
		IncrementQuantity(resp.StatusCode, results)
		m.Unlock()
	}
	wg.Done()
}

func (r *Results) inc() {
	r.Qty++
}

func IncrementQuantity(statusCode int, results *[]Results) {
	fmt.Println("increment", results, statusCode)
	found := false
	for _, result := range *results {
		if result.Status == statusCode {
			fmt.Println("Achou")
			result.inc()
			found = true
			return
		}
	}

	if !found {
		*results = append(*results, Results{
			Status: statusCode,
			Qty:    1,
		})
	}
}

// // docker build -t desafio-tecnico-2 .
// // docker run --name desafio-tecnico-2 desafio-tecnico-2 --foo=123 --blau=334
func main() {
	channel := make(chan int)
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	results := []Results{}

	url := flag.String("url", "", "Url to test")
	concurrency := flag.String("concurrency", "", "Total of concurrent calls")
	requests := flag.String("requests", "", "Total of requests")
	flag.Parse()

	if *url == "" {
		panic(errors.New("url is required"))
	}

	if *concurrency == "" {
		panic(errors.New("concurrency is required"))
	}

	if *requests == "" {
		panic(errors.New("requests is required"))
	}

	concurrencyInt, err := strconv.Atoi(*concurrency)
	if err != nil {
		panic(err)
	}

	requestsInt, err := strconv.Atoi(*requests)
	if err != nil {
		panic(err)
	}

	for i := 0; i < concurrencyInt; i++ {
		fmt.Printf("%s - Create %d\n", time.Now(), i)
		wg.Add(1)
		go someTask(i, *url, channel, &wg, &m, &results)
	}

	for i := 0; i < requestsInt; i++ {
		fmt.Printf("%s - Fill %d\n", time.Now(), i)
		channel <- i
	}

	close(channel)
	wg.Wait()

	for i, result := range results {
		fmt.Printf("%d - Status: %d - Total results: %d\n", i, result.Status, result.Qty)
	}
}
