package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/webdeveloppro/golang-websocket-client/pkg/client"
)

var N_MESSAGES = 5
var N_CLIENTS = 1

var addr = flag.String("addr", ":8000", "http service address")

func worker(id int) {

	client, err := client.NewWebSocketClient(*addr, "frontend")
	if err != nil {
		panic(err)
	}

	fmt.Println("Connecting")

	// write down data every 100 ms
	for i := 0; i < N_MESSAGES; i++ {
		err := client.Write(i)
		if err != nil {
			fmt.Printf("error: %v, writing error\n", err)
		}

		// Sleep to simulate an expensive task.
		time.Sleep(time.Second)
	}

	client.Stop()
}

func main() {

	flag.Parse()

	var wg sync.WaitGroup

	for i := 1; i <= N_CLIENTS; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()

	fmt.Println("Goodbye")
}
