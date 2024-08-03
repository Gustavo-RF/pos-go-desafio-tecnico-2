package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"
)

type Results struct {
	Status int `json:"status"`
	Qty    int `json:"qty"`
}

func callUrl(url string, data chan int, wg *sync.WaitGroup, m *sync.Mutex, results *[]Results) {
	defer wg.Done()
	for _ = range data {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			panic(errors.New("error while formatting address"))
		}

		req.Header.Set("Accepts", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		m.Lock()
		IncrementQuantity(resp.StatusCode, results)
		m.Unlock()
	}
}

func (r *Results) inc() {
	r.Qty++
}

func IncrementQuantity(statusCode int, results *[]Results) {

	found := false
	for i := range *results {
		if (*results)[i].Status == statusCode {
			(*results)[i].inc()
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
	start := time.Now()
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
		wg.Add(1)
		go callUrl(*url, channel, &wg, &m, &results)
	}

	for i := 0; i < requestsInt; i++ {
		channel <- i
	}

	close(channel)
	wg.Wait()

	end := time.Now()

	sort.Slice(results, func(i, j int) bool {
		return results[i].Status < results[j].Status
	})

	fmt.Printf("Stress test finalizado em %s\n", end.Sub(start))
	fmt.Printf("Total de requests realizadas:\t\t%d\n", requestsInt)

	for _, result := range results {
		fmt.Printf("Status %d\t\t\tQtd:\t%d\n", result.Status, result.Qty)
	}
}
