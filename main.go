package main

import (
	"flag"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// func printchannel(c chan int, wg *sync.WaitGroup) {
// 	for i := range c {
// 		fmt.Printf("print %d\n", i)
// 		fmt.Println(i)
// 	}

// 	wg.Done() //notify that we're done here
// }

// func main() {
// 	c := make(chan int)
// 	wg := sync.WaitGroup{}

// 	wg.Add(1) //increase by one to wait for one goroutine to finish
// 	//very important to do it here and not in the goroutine
// 	//otherwise you get race condition

// 	go printchannel(c, &wg) //very important to pass wg by reference
// 	//sync.WaitGroup is a structure, passing it
// 	//by value would produce incorrect results

// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("loop %d\n", i)
// 		c <- i
// 	}

// 	close(c)  //close the channel to terminate the range loop
// 	wg.Wait() //wait for the goroutine to finish
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

func someTask(id int, data chan int, wg *sync.WaitGroup) {
	for taskId := range data {
		// time.Sleep(2 * time.Second)
		fmt.Printf("%s - Worker: %d executed Task %d\n", time.Now(), id, taskId)
	}
	wg.Done()
}

func main() {
	// Creating a channel
	channel := make(chan int)
	wg := sync.WaitGroup{}

	// Define a flag chamada 'foo' com um valor padrão 'default'
	foo := flag.String("foo", "default", "a foo flag")
	blau := flag.String("blau", "default blau", "a blau flag")
	flag.Parse() // Parseia os argumentos de linha de comando

	fmt.Println("FOO:", *foo)
	fmt.Println("BLAU:", *blau)

	fooInt, err := strconv.Atoi(*foo)
	if err != nil {
		// ... handle error
		panic(err)
	}

	blauInt, err := strconv.Atoi(*blau)
	if err != nil {
		// ... handle error
		panic(err)
	}

	// Creating 10.000 workers to execute the task
	for i := 0; i < fooInt; i++ {
		fmt.Printf("%s - Create %d\n", time.Now(), i)
		wg.Add(1)
		go someTask(i, channel, &wg)
	}

	// Filling channel with 100.000 numbers to be executed
	for i := 0; i < blauInt; i++ {
		fmt.Printf("%s - Fill %d\n", time.Now(), i)
		channel <- i
	}

	fmt.Println("Final 1")

	close(channel)
	wg.Wait()
	fmt.Println("Final 2")
}

// package main

// import (
// 	"flag"
// 	"fmt"
// )

// // docker build -t desafio-tecnico-2 .
// // docker run --name desafio-tecnico-2 desafio-tecnico-2 --foo=123 --blau=334
// func main() {
// 	// Define a flag chamada 'foo' com um valor padrão 'default'
// 	foo := flag.String("foo", "default", "a foo flag")
// 	blau := flag.String("blau", "default blau", "a blau flag")
// 	flag.Parse() // Parseia os argumentos de linha de comando

// 	fmt.Println("FOO:", *foo)
// 	fmt.Println("BLAU:", *blau)

// 	ch := make(chan int)
// 	go publish(ch)
// 	reader(ch)
// 	// for x := range ch {
// 	// 	fmt.Printf("Received %d\n", x)
// 	// }
// }

// func reader(ch chan int) {
// 	for x := range ch {
// 		fmt.Printf("Received %d\n", x)
// 	}
// }

// func publish(ch chan int) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("p %d\n", i)
// 		ch <- i
// 	}
// 	close(ch)
// }
